package chat_message

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(chatMessage entity.ChatMessageRepository, id uuid.UUID) (readData *entity.ChatMessage, err error) {
	return chatMessage.Read(id)
}
