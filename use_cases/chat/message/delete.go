package chat_message

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(chatMessage entity.ChatMessageRepository, id uuid.UUID) (err error) {
	return chatMessage.Delete(id)
}
