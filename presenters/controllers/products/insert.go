package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories/product"
	product "github.com/the-mug-codes/service-manager-api/use_cases/product"
)

type insertProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Prices      []*struct {
		Code        string  `json:"code" binding:"required"`
		Title       string  `json:"title"`
		Description string  `json:"description" binding:"required"`
		Value       float32 `json:"value" binding:"required"`
		Status      bool    `json:"status"`
	} `json:"prices"`
	Categories []*uuid.UUID `json:"categories,omitempty"`
	Status     *bool        `json:"status"`
}

func (data insertProduct) dataToInsert() (dataToInsert *entity.Product) {
	dataToInsert = &entity.Product{
		Name:        data.Name,
		Description: data.Description,
		Status:      *data.Status,
	}
	for _, price := range data.Prices {
		dataToInsert.Prices = append(dataToInsert.Prices, entity.ProductPrice{
			Code:        price.Code,
			Title:       price.Title,
			Description: price.Description,
			Value:       price.Value,
			Status:      price.Status,
		})
	}
	for _, category := range data.Categories {
		dataToInsert.Categories = append(dataToInsert.Categories, entity.ProductCategory{
			ID: *category,
		})
	}
	return dataToInsert
}

// @Summary		Create a Product
// @Description	Creates a new product in database.
// @Tags			Products
// @Accept			json
// @Produce		json
// @Param			payload	body		insertProduct	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Product]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/products [post]
func Insert(context *gin.Context) {
	var requestBody *insertProduct
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := product.Insert(repository.Product(context), *requestBody.dataToInsert())
	if err != nil {
		helper.ErrorResponse(context, 401, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
