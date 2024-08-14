package security

import (
	"github.com/gin-gonic/gin"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	twoFactor "github.com/the-mug-codes/adapters-service-api/security/use_cases/two_factor"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Activate Two Factor Authentication for User
// @Description	Activates two factor for validate sign-in.
// @Tags			Security - Two Factor
// @Accept			json
// @Param			code	query		string		true	"Two factor code"
// @Param			id	path		string	true	"Username"
// @Success		200		{object}	helper.ResponseNone
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/auth/two-factor [post]
func TwoFactorActivate(context *gin.Context) {
	username := context.Param("id")
	code := context.Query("code")
	err := twoFactor.Activate(repository.SecurityUser(context), username, code)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot activate two factor", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
