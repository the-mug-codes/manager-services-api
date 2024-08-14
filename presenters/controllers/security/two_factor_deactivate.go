package security

import (
	"github.com/gin-gonic/gin"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	twoFactor "github.com/the-mug-codes/adapters-service-api/security/use_cases/two_factor"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Deactivate Two Factor Authentication for User
// @Description	Deactivates two factor for validate sign-in.
// @Tags			Security - Two Factor
// @Accept			json
// @Param			id	path		string	true	"Username"
// @Success		200		{object}	helper.ResponseNone
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/auth/two-factor [delete]
func TwoFactorDeactivate(context *gin.Context) {
	username := context.Param("id")
	err := twoFactor.Deactivate(repository.SecurityUser(context), username)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot deactivate two factor", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
