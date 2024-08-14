package chat_websocket

import (
	"github.com/google/uuid"
	websocket "github.com/the-mug-codes/service-manager-api/adapters/websocket"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func DeleteChatSection(chatSession entity.ChatSessionRepository, websocketChat *websocket.WebsocketChat, sessionID uuid.UUID) (err error) {
	err = chatSession.Delete(sessionID)
	if err != nil {
		return err
	}
	websocketChat.Sessions.Sessions[sessionID.String()] = &websocket.ChatSession{
		Clients: make(map[string]*websocket.ChatClient),
	}
	client := websocket.CreateChatClientConnection(nil, sessionID, make(chan *entity.ChatMessage))
	websocketChat.Sessions.Unregister <- client
	delete(websocketChat.Sessions.Sessions, sessionID.String())
	close(client.Message)
	return err
}
