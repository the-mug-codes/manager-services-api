package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	repository "github.com/kodit-tecnologia/service-manager/repositories/product"
	product "github.com/kodit-tecnologia/service-manager/use_cases/product"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Show a Products
// @Description	Get a product by idin database.
// @Tags			Products
// @Produce		json
// @Param			id	path		uuid.UUID	true	"ID"
// @Success		200	{object}	helper.ResponseOne[entities.Product]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/products/{id} [get]
func Read(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot read", "id not provided")
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	response, err := product.Read(repository.Product(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
