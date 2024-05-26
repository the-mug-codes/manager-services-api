package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(ticket entity.HelpDeskTicketRepository, dataToInsert entity.HelpDeskTicket) (updatedData *entity.HelpDeskTicket, err error) {
	return ticket.Update(&dataToInsert)
}
