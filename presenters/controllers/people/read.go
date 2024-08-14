package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	person "github.com/the-mug-codes/service-manager-api/use_cases/person"
)

// @Summary		Show a Person
// @Description	Get a person by idin database.
// @Tags			People
// @Produce		json
// @Param			email			query		string		false	"email"
// @Param			phone-number	query		string		false	"phone number"
// @Param			id				path		uuid.UUID	true	"ID"
// @Success		200				{object}	helper.ResponseOne[entities.Person]
// @Failure		400				{object}	helper.Error
// @Failure		401				{object}	helper.Error
// @Failure		404				{object}	helper.Error
// @Router			/people/{id} [get]
func Read(context *gin.Context) {
	email, haveEmail := context.GetQuery("email")
	if haveEmail {
		response, err := person.ReadByEmail(repository.Person(context), email)
		if err != nil {
			helper.ErrorResponse(context, 404, "cannot read", err.Error())
			return
		}
		helper.SuccessResponseOne(context, 200, response)
		return
	}
	phoneNumber, havePhoneNumber := context.GetQuery("phone-number")
	if havePhoneNumber {
		response, err := person.ReadByPhoneNumber(repository.Person(context), phoneNumber)
		if err != nil {
			helper.ErrorResponse(context, 404, "cannot read", err.Error())
			return
		}
		helper.SuccessResponseOne(context, 200, response)
		return
	}
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
	response, err := person.Read(repository.Person(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
