package attachment

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	label "github.com/the-mug-codes/service-manager-api/use_cases/label"
)

// @Summary		Show All Labels
// @Description	Get all labels from database.
// @Tags			Labels
// @Produce		json
// @Success		200	{object}	helper.ResponseMany[[]entities.Label]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/labels [get]
func ReadAll(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := label.ReadAll(repository.Label(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
