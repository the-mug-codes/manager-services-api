package v1

import (
	"github.com/gin-gonic/gin"
	controllerLabels "github.com/kodit-tecnologia/service-manager/presenters/controllers/labels"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Labels(router *gin.RouterGroup) {
	labelsRoute := router.Group("labels")
	{
		labelsRoute.POST("", middleware.Protected(nil, nil), controllerLabels.Insert)
		labelsRoute.GET("", middleware.Protected(nil, nil), controllerLabels.ReadAll)
		labelsRoute.GET(":id", middleware.Protected(nil, nil), controllerLabels.Read)
		labelsRoute.PUT(":id", middleware.Protected(nil, nil), controllerLabels.Update)
		labelsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerLabels.Delete)
	}
}
