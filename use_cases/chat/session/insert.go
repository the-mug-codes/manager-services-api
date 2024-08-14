package chat_session

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(chatSession entity.ChatSessionRepository, dataToInsert entity.ChatSession) (insertedData *entity.ChatSession, err error) {
	return chatSession.Insert(&dataToInsert)
}
