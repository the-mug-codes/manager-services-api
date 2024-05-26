package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(ticket entity.HelpDeskTicketRepository, id uuid.UUID) (readData *entity.HelpDeskTicket, err error) {
	return ticket.Read(id)
}
