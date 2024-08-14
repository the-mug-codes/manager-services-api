package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(ticketMessage entity.HelpDeskTicketMessageRepository, id uuid.UUID) (readData *entity.HelpDeskTicketMessage, err error) {
	return ticketMessage.Read(id)
}
