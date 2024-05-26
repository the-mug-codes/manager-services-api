package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(ticket entity.HelpDeskTicketRepository, page int, pageSize int) (readData *[]entity.HelpDeskTicket, pagination *utils.Pagination, err error) {
	return ticket.ReadAll(page, pageSize)
}
