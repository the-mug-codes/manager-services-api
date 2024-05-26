package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	person "github.com/kodit-tecnologia/service-manager/use_cases/person"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type updatePerson struct {
	Name string `json:"name" binding:"required"` // label name
}

// @Summary		Change a Person
// @Description	Updates a person by idin database.
// @Tags			People
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID		true	"ID"
// @Param			payload	body		updatePerson	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entities.Person]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/people/{id} [put]
func Update(context *gin.Context) {
	var dataToUpdate *entity.Person
	id, haveId := context.Params.Get("id")
	err := context.ShouldBindBodyWith(&dataToUpdate, binding.JSON)
	if err != nil || !haveId {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	dataToUpdate.ID = uuid
	response, err := person.Update(repository.Person(context), *dataToUpdate)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
