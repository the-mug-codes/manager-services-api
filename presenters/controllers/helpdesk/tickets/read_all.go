package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories/helpdesk"
	helpDeskTicket "github.com/the-mug-codes/service-manager-api/use_cases/helpdesk/ticket"
)

// @Summary		Show All HelpDesk Tickets
// @Description	Get all helpdesk tickets from database.
// @Tags			HelpDesk - Tickets
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.HelpDeskTicket]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/helpdesk/tickets [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := helpDeskTicket.ReadAll(repository.HelpDeskTicket(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
