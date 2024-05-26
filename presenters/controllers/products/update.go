package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories/product"
	product "github.com/kodit-tecnologia/service-manager/use_cases/product"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type updateProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Prices      []*struct {
		ID          *uuid.UUID `json:"id"`
		Code        string     `json:"code" binding:"required"`
		Title       string     `json:"title"`
		Description string     `json:"description" binding:"required"`
		Value       float32    `json:"value" binding:"required"`
		Status      bool       `json:"status"`
	} `json:"prices"`
	Categories []*uuid.UUID `json:"categories,omitempty"`
	Status     *bool        `json:"status"`
}

func (data updateProduct) dataToUpdate(id uuid.UUID) (dataToUpdate *entity.Product) {
	dataToUpdate = &entity.Product{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
		Status:      *data.Status,
	}
	for _, price := range data.Prices {
		dataToUpdate.Prices = append(dataToUpdate.Prices, entity.ProductPrice{
			ID:          *price.ID,
			Code:        price.Code,
			Title:       price.Title,
			Description: price.Description,
			Value:       price.Value,
			Status:      price.Status,
		})
	}
	for _, category := range data.Categories {
		dataToUpdate.Categories = append(dataToUpdate.Categories, entity.ProductCategory{
			ID: *category,
		})
	}
	return dataToUpdate
}

// @Summary		Change a Product
// @Description	Updates a product by idin database.
// @Tags			Products
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID		true	"ID"
// @Param			payload	body		updateProduct	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entities.Product]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/products/{id} [put]
func Update(context *gin.Context) {
	var requestBody *updateProduct
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
	response, err := product.Update(repository.Product(context), *requestBody.dataToUpdate(uuid))
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
