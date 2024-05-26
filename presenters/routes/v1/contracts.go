package v1

import (
	"github.com/gin-gonic/gin"
	controllerContracts "github.com/kodit-tecnologia/service-manager/presenters/controllers/contracts"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Contracts(router *gin.RouterGroup) {
	contractsRoute := router.Group("contracts")
	{
		contractsRoute.POST("", middleware.Protected(nil, nil), controllerContracts.Insert)
		contractsRoute.GET("", middleware.Protected(nil, nil), controllerContracts.ReadAll)
		contractsRoute.GET(":id", middleware.Protected(nil, nil), controllerContracts.Read)
		contractsRoute.PUT(":id", middleware.Protected(nil, nil), controllerContracts.Update)
		contractsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerContracts.Delete)
	}
}
