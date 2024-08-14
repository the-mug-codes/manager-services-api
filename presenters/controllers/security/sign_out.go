package security

import (
	"github.com/gin-gonic/gin"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	access "github.com/the-mug-codes/adapters-service-api/security/use_cases/access"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Sign-out User
// @Description	Revokes a new user access credentials.
// @Tags			Security - Auth
// @Accept			json
// @Produce		json
// @Param			id				path		string	true	"Username"
// @Success		200		{object}	helper.ResponseNone
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/auth/sign-out [delete]
func SignOut(context *gin.Context) {
	username := context.Param("id")
	err := access.SignOut(repository.SecurityUser(context), username)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot sign-out user", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
