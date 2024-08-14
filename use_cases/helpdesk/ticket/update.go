package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(ticket entity.HelpDeskTicketRepository, dataToInsert entity.HelpDeskTicket) (updatedData *entity.HelpDeskTicket, err error) {
	return ticket.Update(&dataToInsert)
}
