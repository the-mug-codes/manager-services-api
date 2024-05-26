package attachment

import (
	"github.com/gin-gonic/gin"
	repository "github.com/kodit-tecnologia/service-manager/repositories/helpdesk"
	helpDeskTicket "github.com/kodit-tecnologia/service-manager/use_cases/helpdesk/ticket"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
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
