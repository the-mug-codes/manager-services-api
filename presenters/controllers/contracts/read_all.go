package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	contract "github.com/the-mug-codes/service-manager-api/use_cases/contract"
)

// @Summary		Show All Contract
// @Description	Get all contracts from database.
// @Tags			Contracts
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.Contract]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/contracts [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := contract.ReadAll(repository.Contract(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
