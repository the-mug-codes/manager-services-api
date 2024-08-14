package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerPeople "github.com/the-mug-codes/service-manager-api/presenters/controllers/people"
)

func People(router *gin.RouterGroup) {
	peopleRoute := router.Group("people")
	{
		peopleRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerPeople.Insert)
		peopleRoute.GET("", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerPeople.ReadAll)
		peopleRoute.GET(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerPeople.Read)
		peopleRoute.PUT(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerPeople.Update)
		peopleRoute.DELETE(":id", middleware.Protected(&[]string{"admin:full"}), controllerPeople.Delete)
	}
}
