package chat_message

import (
	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAllBySession(chatMessage entity.ChatMessageRepository, sessionID uuid.UUID, page int, pageSize int) (readData *[]entity.ChatMessage, pagination *utils.Pagination, err error) {
	return chatMessage.ReadAllBySession(sessionID, page, pageSize)
}
