package chat_session

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(chatSession entity.ChatSessionRepository, id uuid.UUID) (readData *entity.ChatSession, err error) {
	return chatSession.Read(id)
}
