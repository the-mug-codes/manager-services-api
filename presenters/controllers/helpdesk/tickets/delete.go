package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories/helpdesk"
	helpDeskTicket "github.com/the-mug-codes/service-manager-api/use_cases/helpdesk/ticket"
)

// @Summary		Delete a HelpDesk Ticket
// @Description	Removes a helpdesk ticket by idfrom database.
// @Tags			HelpDesk - Tickets
// @Produce		json
// @Param			id	path		uuid.UUID	true	"ID"
// @Success		200	{object}	helper.ResponseNone
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/helpdesk/tickets/{id} [delete]
func Delete(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot bind data", "resource id not provided")
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot delete", err.Error())
		return
	}
	err = helpDeskTicket.Delete(repository.HelpDeskTicket(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot delete", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
