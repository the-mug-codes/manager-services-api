package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(ticketMessage entity.HelpDeskTicketMessageRepository, id uuid.UUID) (readData *entity.HelpDeskTicketMessage, err error) {
	return ticketMessage.Read(id)
}
