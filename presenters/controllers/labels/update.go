package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	label "github.com/the-mug-codes/service-manager-api/use_cases/label"
)

type updateLabel struct {
	Name string `json:"name" binding:"required"`
}

func (data updateLabel) dataToUpdate(id uuid.UUID) (dataToUpdate *entity.Label) {
	dataToUpdate = &entity.Label{
		ID:   id,
		Name: data.Name,
	}
	return dataToUpdate
}

// @Summary		Change a Label
// @Description	Updates a label by idin database.
// @Tags			Labels
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID	true	"ID"
// @Param			payload	body		updateLabel	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entities.Label]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/labels/{id} [put]
func Update(context *gin.Context) {
	var requestBody *updateLabel
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
	response, err := label.Update(repository.Label(context), *requestBody.dataToUpdate(uuid))
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
