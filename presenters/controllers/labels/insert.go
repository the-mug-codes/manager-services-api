package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	label "github.com/the-mug-codes/service-manager-api/use_cases/label"
)

type insertLabel struct {
	Name string `json:"name" binding:"required"`
}

func (data insertLabel) dataToInsert() (dataToInsert *entity.Label) {
	return &entity.Label{
		Name: data.Name,
	}
}

// @Summary		Create a Label
// @Description	Creates a new label in database.
// @Tags			Labels
// @Accept			json
// @Produce		json
// @Param			payload	body		insertLabel	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Label]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/labels [post]
func Insert(context *gin.Context) {
	var requestBody *insertLabel
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := label.Insert(repository.Label(context), *requestBody.dataToInsert())
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
