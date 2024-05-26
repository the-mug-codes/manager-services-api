package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(ticketMessage entity.HelpDeskTicketMessageRepository, dataToInsert entity.HelpDeskTicketMessage) (insertedData *entity.HelpDeskTicketMessage, err error) {
	return ticketMessage.Insert(&dataToInsert)
}
