package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	repository "github.com/kodit-tecnologia/service-manager/repositories/helpdesk"
	helpDeskTicket "github.com/kodit-tecnologia/service-manager/use_cases/helpdesk/ticket"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Show a HelpDesk Ticket
// @Description	Get a helpdesk ticket by idin database.
// @Tags			HelpDesk - Tickets
// @Produce		json
// @Param			id	path		uuid.UUID	true	"ID"
// @Success		200	{object}	helper.ResponseOne[entities.HelpDeskTicket]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/helpdesk/tickets/{id} [get]
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
	response, err := helpDeskTicket.Read(repository.HelpDeskTicket(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
