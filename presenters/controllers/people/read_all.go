package attachment

import (
	"github.com/gin-gonic/gin"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	person "github.com/kodit-tecnologia/service-manager/use_cases/person"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
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
