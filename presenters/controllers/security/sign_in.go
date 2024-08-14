package security

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	entity "github.com/the-mug-codes/adapters-service-api/security"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	access "github.com/the-mug-codes/adapters-service-api/security/use_cases/access"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Sign-in User
// @Description	Creates a new user access credentials.
// @Tags			Security - Auth
// @Accept			json
// @Produce		json
// @Param			code	query		string		false	"Two factor code"
// @Param			google_access_token	query		string		false	"Google access token"
// @Param			payload	body		entity.SignIn	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entity.Token]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/auth/sign-in [post]
func SignIn(context *gin.Context) {
	code := context.Query("code")
	googleAccessToken := context.Query("google_access_token")
	var signInData *entity.SignIn
	err := context.ShouldBindBodyWith(&signInData, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind json", err.Error())
		return
	}
	response, err := access.SignIn(repository.SecurityUser(context), signInData, code, googleAccessToken)
	if err != nil {
		helper.ErrorResponse(context, 401, "cannot sign-in user", err.Error())
		return
	}
	helper.SuccessResponseOne[entity.Token](context, 200, &response)
}
