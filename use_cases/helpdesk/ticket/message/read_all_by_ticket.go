package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAllByTicket(ticketMessage entity.HelpDeskTicketMessageRepository, ticketID uuid.UUID, page int, pageSize int) (readData *[]entity.HelpDeskTicketMessage, pagination *utils.Pagination, err error) {
	return ticketMessage.ReadAllByTicket(ticketID, page, pageSize)
}
