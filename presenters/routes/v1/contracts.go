package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerContracts "github.com/the-mug-codes/service-manager-api/presenters/controllers/contracts"
)

func Contracts(router *gin.RouterGroup) {
	contractsRoute := router.Group("contracts")
	{
		contractsRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerContracts.Insert)
		contractsRoute.GET("", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerContracts.ReadAll)
		contractsRoute.GET(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerContracts.Read)
		contractsRoute.PUT(":id", middleware.Protected(&[]string{"admin:full"}), controllerContracts.Update)
		contractsRoute.DELETE(":id", middleware.Protected(&[]string{"admin:full"}), controllerContracts.Delete)
	}
}
