package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories/product"
	productCategory "github.com/the-mug-codes/service-manager-api/use_cases/product/category"
)

// @Summary		Show All Products Categories
// @Description	Get all products categories from database.
// @Tags			Products - Categories
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.ProductCategory]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/products/categories [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := productCategory.ReadAll(repository.ProductCategory(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
