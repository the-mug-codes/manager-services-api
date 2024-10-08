package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	person "github.com/the-mug-codes/service-manager-api/use_cases/person"
)

type insertPerson struct {
	Name string `json:"name" binding:"required"` // label name
}

// @Summary		Create a Person
// @Description	Creates a new person in database.
// @Tags			People
// @Accept			json
// @Produce		json
// @Param			payload	body		insertPerson	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Person]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/people [post]
func Insert(context *gin.Context) {
	var dataToInsert *entity.Person
	err := context.ShouldBindBodyWith(&dataToInsert, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := person.Insert(repository.Person(context), *dataToInsert)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
