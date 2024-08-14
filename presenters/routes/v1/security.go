package v1

import (
	"github.com/gin-gonic/gin"
	controllerSecurity "github.com/the-mug-codes/service-manager-api/presenters/controllers/security"
)

func Security(router *gin.RouterGroup) {
	authRoute := router.Group("auth")
	{
		authRoute.POST("sign-in", controllerSecurity.SignIn)
		authRoute.DELETE("sign-out", controllerSecurity.SignOut)
		tokenRoute := authRoute.Group("token")
		{
			tokenRoute.PUT("refresh", controllerSecurity.RefreshToken)
		}
		googleRoute := authRoute.Group("google")
		{
			googleRoute.POST(":id", controllerSecurity.GoogleActivate)
			googleRoute.DELETE(":id", controllerSecurity.GoogleDeactivate)
		}
		twoFactorRoute := authRoute.Group("two-factor")
		{
			twoFactorRoute.GET(":id", controllerSecurity.TwoFactorGenerate)
			twoFactorRoute.POST(":id", controllerSecurity.TwoFactorActivate)
			twoFactorRoute.DELETE(":id", controllerSecurity.TwoFactorDeactivate)
		}
	}
	passwordRoute := router.Group("password")
	{
		passwordRoute.POST("forgot/:id", controllerSecurity.PasswordForgot)
		passwordRoute.POST("change", controllerSecurity.PasswordChange)
	}
}
