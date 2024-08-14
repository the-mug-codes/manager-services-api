package people

import (
	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAllByTicket(ticketMessage entity.HelpDeskTicketMessageRepository, ticketID uuid.UUID, page int, pageSize int) (readData *[]entity.HelpDeskTicketMessage, pagination *utils.Pagination, err error) {
	return ticketMessage.ReadAllByTicket(ticketID, page, pageSize)
}
