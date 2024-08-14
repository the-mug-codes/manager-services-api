package chat_message

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(chatMessage entity.ChatMessageRepository, page int, pageSize int) (readData *[]entity.ChatMessage, pagination *utils.Pagination, err error) {
	return chatMessage.ReadAll(page, pageSize)
}
