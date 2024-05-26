package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	contract "github.com/kodit-tecnologia/service-manager/use_cases/contract"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Show a Contract
// @Description	Get a contract by idin database.
// @Tags			Contracts
// @Produce		json
// @Param			id	path		uuid.UUID	true	"ID"
// @Success		200	{object}	helper.ResponseOne[entities.Contract]
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/contracts/{id} [get]
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
	response, err := contract.Read(repository.Contract(context), uuid)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
