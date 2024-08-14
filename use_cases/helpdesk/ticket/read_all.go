package people

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(ticket entity.HelpDeskTicketRepository, page int, pageSize int) (readData *[]entity.HelpDeskTicket, pagination *utils.Pagination, err error) {
	return ticket.ReadAll(page, pageSize)
}
