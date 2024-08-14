package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(ticketMessage entity.HelpDeskTicketMessageRepository, dataToInsert entity.HelpDeskTicketMessage) (insertedData *entity.HelpDeskTicketMessage, err error) {
	return ticketMessage.Insert(&dataToInsert)
}
