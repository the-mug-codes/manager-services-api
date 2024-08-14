package security

import (
	"github.com/gin-gonic/gin"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	twoFactor "github.com/the-mug-codes/adapters-service-api/security/use_cases/two_factor"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Generate Two Factor Authentication for User
// @Description	Generate two factor for validate sign-in.
// @Tags			Security - Two Factor
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Username"
// @Success		201		{object}	helper.ResponseOne[twoFactor.OtpResponse]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/auth/two-factor [get]
func TwoFactorGenerate(context *gin.Context) {
	username := context.Param("id")
	response, err := twoFactor.Generate(repository.SecurityUser(context), username)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot create two factor", err.Error())
		return
	}
	helper.SuccessResponseOne[twoFactor.OtpResponse](context, 201, &response)
}
