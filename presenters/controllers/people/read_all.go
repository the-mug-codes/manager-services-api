package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	person "github.com/the-mug-codes/service-manager-api/use_cases/person"
)

// @Summary		Show All People
// @Description	Get all people from database.
// @Tags			People
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.Person]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/people [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := person.ReadAll(repository.Person(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
