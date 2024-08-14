package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories/helpdesk"
	helpDeskTicketMessage "github.com/the-mug-codes/service-manager-api/use_cases/helpdesk/ticket/message"
)

// @Summary		Show All HelpDesk Tickets Messages
// @Description	Get all helpdesk tickets messages from database.
// @Tags			HelpDesk - Tickets Messages
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.HelpDeskTicketMessage]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/helpdesk/tickets/messages [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := helpDeskTicketMessage.ReadAll(repository.HelpDeskTicketMessage(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
