package security

import (
	"github.com/gin-gonic/gin"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	google "github.com/the-mug-codes/adapters-service-api/security/use_cases/google"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Deactivate Google Sign-in for User
// @Description	Deactivates sign-in with a Google account.
// @Tags			Security - Google
// @Accept			json
// @Param			id	path		string	true	"Username"
// @Success		200		{object}	helper.ResponseNone
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/auth/google [delete]
func GoogleDeactivate(context *gin.Context) {
	username := context.Param("id")
	err := google.Deactivate(repository.SecurityUser(context), username)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot deactivate google account sign-in", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
