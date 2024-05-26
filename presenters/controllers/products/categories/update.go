package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories/product"
	productCategory "github.com/kodit-tecnologia/service-manager/use_cases/product/category"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type updateProductCategory struct {
	Kind        string       `json:"kind" binding:"required"`
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Products    []*uuid.UUID ` json:"products,omitempty"`
	Status      *bool        ` json:"status"`
}

func (data updateProductCategory) dataToUpdate(id uuid.UUID) (dataToUpdate *entity.ProductCategory) {
	dataToUpdate = &entity.ProductCategory{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
		Status:      *data.Status,
	}
	for _, product := range data.Products {
		dataToUpdate.Products = append(dataToUpdate.Products, entity.Product{
			ID: *product,
		})
	}
	return dataToUpdate
}

// @Summary		Change a Product Category
// @Description	Updates a product category by idin database.
// @Tags			Products - Categories
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID				true	"ID"
// @Param			payload	body		updateProductCategory	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entities.ProductCategory]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/products/categories/{id} [put]
func Update(context *gin.Context) {
	var requestBody *updateProductCategory
	id, haveId := context.Params.Get("id")
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil || !haveId {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	response, err := productCategory.Update(repository.ProductCategory(context), *requestBody.dataToUpdate(uuid))
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
