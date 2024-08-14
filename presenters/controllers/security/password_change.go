package security

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	entity "github.com/the-mug-codes/adapters-service-api/security"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	recovery "github.com/the-mug-codes/adapters-service-api/security/use_cases/recovery"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Change a User Password
// @Description	Changes a user password.
// @Tags			Security - Password
// @Accept			json
// @Produce		json
// @Param			payload	body		entity.PasswordRecovery	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entity.Token]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/password/change [post]
func PasswordChange(context *gin.Context) {
	var recoveryPasswordData *entity.PasswordRecovery
	err := context.ShouldBindBodyWith(&recoveryPasswordData, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind json", err.Error())
		return
	}
	err = recovery.ChangePassword(repository.SecurityUser(context), recoveryPasswordData)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot change password", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
