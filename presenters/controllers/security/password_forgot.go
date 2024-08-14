package security

import (
	"github.com/gin-gonic/gin"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	recovery "github.com/the-mug-codes/adapters-service-api/security/use_cases/recovery"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Set a User Password as Forgot
// @Description	Sets a user password as forgotten.
// @Tags			Security - Password
// @Accept			json
// @Param			id	path		string	true	"Username"
// @Success		200		{object}	helper.ResponseNone
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/password/forgot [post]
func PasswordForgot(context *gin.Context) {
	username := context.Param("id")
	err := recovery.ForgotPassword(repository.SecurityUser(context), username)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot set password as forgotten", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
