package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(ticket entity.HelpDeskTicketRepository, dataToInsert entity.HelpDeskTicket) (insertedData *entity.HelpDeskTicket, err error) {
	return ticket.Insert(&dataToInsert)
}
