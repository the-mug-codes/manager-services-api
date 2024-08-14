package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerLabels "github.com/the-mug-codes/service-manager-api/presenters/controllers/labels"
)

func Labels(router *gin.RouterGroup) {
	labelsRoute := router.Group("labels")
	{
		labelsRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerLabels.Insert)
		labelsRoute.GET("", middleware.Protected(&[]string{"admin:full"}), controllerLabels.ReadAll)
		labelsRoute.GET(":id", middleware.Protected(&[]string{"admin:full"}), controllerLabels.Read)
		labelsRoute.PUT(":id", middleware.Protected(&[]string{"admin:full"}), controllerLabels.Update)
		labelsRoute.DELETE(":id", middleware.Protected(&[]string{"admin:full"}), controllerLabels.Delete)
	}
}
