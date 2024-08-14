package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories/product"

	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	productCategory "github.com/the-mug-codes/service-manager-api/use_cases/product/category"
)

type insertProductCategory struct {
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Products    []*uuid.UUID ` json:"products,omitempty"`
	Status      *bool        ` json:"status"`
}

func (data insertProductCategory) dataToInsert() (dataToInsert *entity.ProductCategory) {
	dataToInsert = &entity.ProductCategory{
		Name:        data.Name,
		Description: data.Description,
		Status:      *data.Status,
	}
	for _, product := range data.Products {
		dataToInsert.Products = append(dataToInsert.Products, entity.Product{
			ID: *product,
		})
	}
	return dataToInsert
}

// @Summary		Create a Product Category
// @Description	Creates a new product category in database.
// @Tags			Products - Categories
// @Accept			json
// @Produce		json
// @Param			payload	body		insertProductCategory	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.ProductCategory]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/products/categories [post]
func Insert(context *gin.Context) {
	var requestBody *insertProductCategory
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := productCategory.Insert(repository.ProductCategory(context), *requestBody.dataToInsert())
	if err != nil {
		helper.ErrorResponse(context, 401, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
