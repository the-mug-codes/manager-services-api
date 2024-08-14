package people

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(ticketMessage entity.HelpDeskTicketMessageRepository, page int, pageSize int) (readData *[]entity.HelpDeskTicketMessage, pagination *utils.Pagination, err error) {
	return ticketMessage.ReadAll(page, pageSize)
}
