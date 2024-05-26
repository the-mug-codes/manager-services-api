package v1

import (
	"github.com/gin-gonic/gin"
	controllerLabels "github.com/kodit-tecnologia/service-manager/presenters/controllers/labels"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Labels(router *gin.RouterGroup) {
	labelsRoute := router.Group("labels")
	{
		labelsRoute.POST("", controllerLabels.Insert)
		labelsRoute.GET("", controllerLabels.ReadAll)
		labelsRoute.GET(":id", controllerLabels.Read)
		labelsRoute.PUT(":id", controllerLabels.Update)
		labelsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerLabels.Delete)
	}
}
