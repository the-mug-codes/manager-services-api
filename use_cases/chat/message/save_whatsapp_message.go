package chat_message

import (
	"strconv"

	"github.com/the-mug-codes/service-manager-api/adapters/websocket"
	"github.com/the-mug-codes/service-manager-api/adapters/whatsapp"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func SaveWhatsAppMessage(chatSession entity.ChatSessionRepository, chatMessage entity.ChatMessageRepository, websocketChat *websocket.WebsocketChat, message *whatsapp.MessageParsed) (err error) {
	var session *entity.ChatSession
	phoneNumber, err := strconv.Atoi(*message.ContactPhoneNumber)
	if err != nil {
		return err
	}
	session, _ = chatSession.ReadByPhoneNumber(phoneNumber)
	if session == nil || !session.Status {
		newChatSession := &entity.ChatSession{
			Channel:     "whatsapp",
			PhoneNumber: &phoneNumber,
		}
		session, err = chatSession.Insert(newChatSession)
		if err != nil {
			return err
		}
		websocketChat.Sessions.Sessions[string(session.ID.String())] = &websocket.ChatSession{
			ID:      session.ID.String(),
			Clients: make(map[string]*websocket.ChatClient),
		}
	}
	newChatMessage := &entity.ChatMessage{
		SessionID:   session.ID,
		Name:        *message.ContactDisplayName,
		MessageType: *message.MessageType,
		Body:        message.Message,
		MediaID:     message.MediaID,
		ActionID:    message.ActionID,
	}
	_, err = chatMessage.Insert(newChatMessage)
	if err != nil {
		return err
	}
	websocketChat.Sessions.Broadcast <- newChatMessage
	return err
}
