package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(ticketMessage entity.HelpDeskTicketMessageRepository, page int, pageSize int) (readData *[]entity.HelpDeskTicketMessage, pagination *utils.Pagination, err error) {
	return ticketMessage.ReadAll(page, pageSize)
}
