package security

import (
	"github.com/gin-gonic/gin"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	google "github.com/the-mug-codes/adapters-service-api/security/use_cases/google"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Activate Google Sign-in for User
// @Description	Activates sign-in with a Google account.
// @Tags			Security - Google
// @Accept			json
// @Param			google_access_token	query		string		true	"Google access token"
// @Param			id	path		string	true	"Username"
// @Success		201		{object}	helper.ResponseNone
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/auth/google [post]
func GoogleActivate(context *gin.Context) {
	username := context.Param("id")
	googleAccessToken := context.Query("google_access_token")
	err := google.Activate(repository.SecurityUser(context), googleAccessToken, username)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot activate google account sign-in", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 201)
}
