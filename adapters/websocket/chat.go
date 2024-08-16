package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/the-mug-codes/service-manager-api/adapters/whatsapp"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

type WebsocketChatSessions struct {
	Sessions   map[string]*ChatSession
	Register   chan *ChatClient
	Unregister chan *ChatClient
	Broadcast  chan *entity.ChatMessage
}

type WebsocketChat struct {
	Sessions *WebsocketChatSessions
}

var websocketChat *WebsocketChat

func StartWebsocketChat() {
	websocketChat = &WebsocketChat{
		Sessions: &WebsocketChatSessions{
			Sessions:   make(map[string]*ChatSession),
			Register:   make(chan *ChatClient),
			Unregister: make(chan *ChatClient),
			Broadcast:  make(chan *entity.ChatMessage, 5),
		},
	}
	log.Print("\033[1m\033[7m\033[32m online \033[0m \033[1mwebsocket - live chat\033[0m")
	go websocketChat.Sessions.Run()
}

func GetWebsocketChat() *WebsocketChat {
	return websocketChat
}

func (chat *WebsocketChatSessions) Run() {
	for {
		select {
		case client := <-chat.Register:
			if _, ok := chat.Sessions[client.SessionID.String()]; ok {
				session := chat.Sessions[client.SessionID.String()]
				if _, ok := session.Clients[client.ID.String()]; !ok {
					session.Clients[client.ID.String()] = client
				}
			}
		case client := <-chat.Unregister:
			if _, ok := chat.Sessions[client.SessionID.String()]; ok {
				session := chat.Sessions[client.SessionID.String()]
				if _, ok := session.Clients[client.ID.String()]; ok {
					delete(session.Clients, client.ID.String())
					close(client.Message)
				}
			}
		case message := <-chat.Broadcast:
			if _, ok := chat.Sessions[message.SessionID.String()]; ok {
				for _, client := range chat.Sessions[message.SessionID.String()].Clients {
					client.Message <- message
				}
			}
		}
	}
}

type ChatClient struct {
	Conn      *websocket.Conn
	ID        uuid.UUID
	SessionID uuid.UUID
	Message   chan *entity.ChatMessage
}

func (chatClient *ChatClient) WriteMessage() {
	defer func() {
		chatClient.Conn.Close()
	}()
	for {
		message, ok := <-chatClient.Message
		if !ok {
			return
		}
		err := chatClient.Conn.WriteJSON(message)
		if err != nil {
			log.Println("Error writing message:", err)
			return
		}
	}
}

func (chatClient *ChatClient) ReadMessage(chatHubSessions *WebsocketChatSessions, chatSession entity.ChatSessionRepository, chatMessage entity.ChatMessageRepository, whatsappConnection whatsapp.WhatsappInterface) {
	defer func() {
		chatHubSessions.Unregister <- chatClient
		chatClient.Conn.Close()
	}()
	for {
		_, message, err := chatClient.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Unexpected close error:", err)
			} else {
				log.Println("Error reading message:", err)
			}
			return
		}
		var msg entity.ChatMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Unexpected message error:", err)
		}
		session, err := chatSession.Read(msg.SessionID)
		if err != nil {
			log.Println("Error on reading session:", err)
			return
		}
		if session.Channel == "whatsapp" {
			err = whatsappConnection.SendTextMessage(fmt.Sprintf("%v", *session.PhoneNumber), whatsapp.SendTextMessage{
				Body:       *msg.Body,
				PreviewUrl: true,
			})
			if err != nil {
				log.Println("Error on sending whatsapp message:", err)
				return
			}
		}
		_, err = chatMessage.Insert(&msg)
		if err != nil {
			log.Println("Error on saving message:", err)
			return
		}
		chatHubSessions.Broadcast <- &msg
	}
}

func CreateChatClientConnection(conn *websocket.Conn, sessionID uuid.UUID, message chan *entity.ChatMessage) *ChatClient {
	return &ChatClient{
		Conn:      conn,
		ID:        uuid.New(),
		SessionID: sessionID,
		Message:   message,
	}
}

type ChatSession struct {
	ID      string                 `json:"id" binding:"required"`
	Clients map[string]*ChatClient `json:"clients"`
}
