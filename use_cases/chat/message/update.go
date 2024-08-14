package chat_message

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(chatMessage entity.ChatMessageRepository, dataToInsert entity.ChatMessage) (updatedData *entity.ChatMessage, err error) {
	return chatMessage.Update(&dataToInsert)
}
