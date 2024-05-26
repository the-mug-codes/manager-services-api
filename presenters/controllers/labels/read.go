package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	label "github.com/kodit-tecnologia/service-manager/use_cases/label"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Show a Label
// @Description	Get a label by idin database.
// @Tags			Labels
// @Produce		json
// @Param			id	path		uuid.UUID	true	"ID"
// @Success		200	{object}	helper.ResponseOne[entities.Label]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/labels/{id} [get]
func Read(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 404, "cannot read", "id not provided")
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	response, err := label.Read(repository.Label(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
