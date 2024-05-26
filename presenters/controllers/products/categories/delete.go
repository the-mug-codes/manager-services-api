package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	repository "github.com/kodit-tecnologia/service-manager/repositories/product"
	productCategory "github.com/kodit-tecnologia/service-manager/use_cases/product/category"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Delete a Product Category
// @Description	Removes a product category by idfrom database.
// @Tags			Products - Categories
// @Produce		json
// @Param			id	path		uuid.UUID	true	"ID"
// @Success		200	{object}	helper.ResponseNone
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/products/categories/{id} [delete]
func Delete(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot bind data", "resource id not provided")
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot delete", err.Error())
		return
	}
	err = productCategory.Delete(repository.ProductCategory(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot delete", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
