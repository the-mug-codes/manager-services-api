package v1

import (
	"github.com/gin-gonic/gin"
	controllerContracts "github.com/kodit-tecnologia/service-manager/presenters/controllers/contracts"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Contracts(router *gin.RouterGroup) {
	contractsRoute := router.Group("contracts")
	{
		contractsRoute.POST("", controllerContracts.Insert)
		contractsRoute.GET("", controllerContracts.ReadAll)
		contractsRoute.GET(":id", controllerContracts.Read)
		contractsRoute.PUT(":id", controllerContracts.Update)
		contractsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerContracts.Delete)
	}
}
