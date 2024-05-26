package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(ticket entity.HelpDeskTicketRepository, dataToInsert entity.HelpDeskTicket) (insertedData *entity.HelpDeskTicket, err error) {
	return ticket.Insert(&dataToInsert)
}
