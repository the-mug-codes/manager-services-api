package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	project "github.com/the-mug-codes/service-manager-api/use_cases/project"
)

// @Summary		Show All Projects
// @Description	Get all projects from database.
// @Tags			Projects
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.Project]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/projects [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := project.ReadAll(repository.Project(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
