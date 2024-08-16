package chat_websocket

import (
	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	websocket "github.com/the-mug-codes/service-manager-api/adapters/websocket"
	whatsapp "github.com/the-mug-codes/service-manager-api/adapters/whatsapp"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func JoinChatSection(chatSession entity.ChatSessionRepository, chatMessage entity.ChatMessageRepository, whatsappConnection whatsapp.WhatsappInterface, websocketChat *websocket.WebsocketChat, wsConnection *ws.Conn, sessionID uuid.UUID) (err error) {
	if _, ok := websocketChat.Sessions.Sessions[sessionID.String()]; !ok {
		websocketChat.Sessions.Sessions[sessionID.String()] = &websocket.ChatSession{
			ID:      sessionID.String(),
			Clients: make(map[string]*websocket.ChatClient),
		}
	}
	client := websocket.CreateChatClientConnection(wsConnection, sessionID, make(chan *entity.ChatMessage))
	websocketChat.Sessions.Register <- client
	go client.WriteMessage()
	client.ReadMessage(websocketChat.Sessions, chatSession, chatMessage, whatsappConnection)
	return err
}
