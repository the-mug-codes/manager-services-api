package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	invoice "github.com/the-mug-codes/service-manager-api/use_cases/invoice"
)

// @Summary		Show All Invoices
// @Description	Get all invoices from database.
// @Tags			Invoices
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.Invoice]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/invoices [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := invoice.ReadAll(repository.Invoice(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
