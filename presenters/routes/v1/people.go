package v1

import (
	"github.com/gin-gonic/gin"
	controllerPeople "github.com/kodit-tecnologia/service-manager/presenters/controllers/people"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func People(router *gin.RouterGroup) {
	peopleRoute := router.Group("people")
	{
		peopleRoute.POST("", middleware.Protected(nil, nil), controllerPeople.Insert)
		peopleRoute.GET("", middleware.Protected(nil, nil), controllerPeople.ReadAll)
		peopleRoute.GET(":id", middleware.Protected(nil, nil), controllerPeople.Read)
		peopleRoute.PUT(":id", middleware.Protected(nil, nil), controllerPeople.Update)
		peopleRoute.DELETE(":id", middleware.Protected(nil, nil), controllerPeople.Delete)
	}
}
