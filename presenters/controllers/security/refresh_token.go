package security

import (
	"github.com/gin-gonic/gin"
	entity "github.com/the-mug-codes/adapters-service-api/security"
	repository "github.com/the-mug-codes/adapters-service-api/security"
	token "github.com/the-mug-codes/adapters-service-api/security/use_cases/token"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Refesh Tokens
// @Description	Refresh user access tokens.
// @Tags			Security - Token
// @Accept			json
// @Produce		json
// @Success		201		{object}	helper.ResponseOne[entity.Token]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/token/refresh [put]
func RefreshToken(context *gin.Context) {
	authorization := context.GetHeader("Authorization")
	if len(authorization) <= 0 {
		helper.ErrorResponse(context, 401, "cannot refresh token", "authorization header not provided")
		return
	}
	response, err := token.Refresh(repository.SecurityUser(context), helper.GetTokenHash(authorization))
	if err != nil {
		helper.ErrorResponse(context, 401, "cannot refresh token", err.Error())
		return
	}
	helper.SuccessResponseOne[entity.Token](context, 201, &response)
}
