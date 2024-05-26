package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Delete(ticket entity.HelpDeskTicketRepository, id uuid.UUID) (err error) {
	return ticket.Delete(id)
}
