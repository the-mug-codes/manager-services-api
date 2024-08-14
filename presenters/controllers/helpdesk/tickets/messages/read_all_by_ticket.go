package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories/helpdesk"
	helpDeskTicketMessage "github.com/the-mug-codes/service-manager-api/use_cases/helpdesk/ticket/message"
)

// @Summary		Show All HelpDesk Ticket Messages
// @Description	Get all messages from a helpdesk ticket from database.
// @Tags			HelpDesk - Tickets Messages
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.HelpDeskTicketMessage]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/helpdesk/tickets/:id/messages [get]
func ReadAllByTicket(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 404, "cannot read", "id not provided")
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := helpDeskTicketMessage.ReadAllByTicket(repository.HelpDeskTicketMessage(context), uuid, page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
