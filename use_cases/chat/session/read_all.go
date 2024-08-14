package chat_session

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(chatSession entity.ChatSessionRepository, page int, pageSize int) (readData *[]entity.ChatSession, pagination *utils.Pagination, err error) {
	return chatSession.ReadAll(page, pageSize)
}
