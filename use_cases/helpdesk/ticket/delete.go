package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(ticket entity.HelpDeskTicketRepository, id uuid.UUID) (err error) {
	return ticket.Delete(id)
}
