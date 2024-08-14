package chat_message

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(chatMessage entity.ChatMessageRepository, dataToInsert entity.ChatMessage) (insertedData *entity.ChatMessage, err error) {
	return chatMessage.Insert(&dataToInsert)
}
