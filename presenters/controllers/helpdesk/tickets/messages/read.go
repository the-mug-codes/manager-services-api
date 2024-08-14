package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories/helpdesk"
	helpDeskTicketMessage "github.com/the-mug-codes/service-manager-api/use_cases/helpdesk/ticket/message"
)

// @Summary		Show a HelpDesk Ticket Message
// @Description	Get a helpdesk ticket message by idin database.
// @Tags			HelpDesk - Tickets Messages
// @Produce		json
// @Param			id	path		uuid.UUID	true	"ID"
// @Success		200	{object}	helper.ResponseOne[entities.HelpDeskTicketMessage]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/helpdesk/tickets/messages/{id} [get]
func Read(context *gin.Context) {
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
	response, err := helpDeskTicketMessage.Read(repository.HelpDeskTicketMessage(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
