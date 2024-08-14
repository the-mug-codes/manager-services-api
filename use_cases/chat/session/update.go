package chat_session

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(chatSession entity.ChatSessionRepository, dataToInsert entity.ChatSession) (updatedData *entity.ChatSession, err error) {
	return chatSession.Update(&dataToInsert)
}
