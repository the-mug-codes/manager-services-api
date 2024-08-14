package chat_session

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(chatSession entity.ChatSessionRepository, id uuid.UUID) (err error) {
	return chatSession.Delete(id)
}
