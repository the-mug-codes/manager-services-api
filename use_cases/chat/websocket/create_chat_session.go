package chat_websocket

import (
	"github.com/google/uuid"
	websocket "github.com/the-mug-codes/service-manager-api/adapters/websocket"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func CreateChatSection(chatSession entity.ChatSessionRepository, websocketChat *websocket.WebsocketChat, email *string, phoneNumber *int) (sessionID uuid.UUID, err error) {
	newChatSession := &entity.ChatSession{
		Channel:     "live-chat",
		PhoneNumber: phoneNumber,
		Email:       email,
	}
	session, err := chatSession.Insert(newChatSession)
	if err != nil {
		return sessionID, err
	}
	websocketChat.Sessions.Sessions[session.ID.String()] = &websocket.ChatSession{
		ID:      sessionID.String(),
		Clients: make(map[string]*websocket.ChatClient),
	}
	return session.ID, err
}
