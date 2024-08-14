package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerProjects "github.com/the-mug-codes/service-manager-api/presenters/controllers/projects"
)

func Projects(router *gin.RouterGroup) {
	projectsRoute := router.Group("projects")
	{
		projectsRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerProjects.Insert)
		projectsRoute.GET("", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerProjects.ReadAll)
		projectsRoute.GET(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerProjects.Read)
		projectsRoute.PUT(":id", middleware.Protected(&[]string{"admin:full"}), controllerProjects.Update)
		projectsRoute.DELETE(":id", middleware.Protected(&[]string{"admin:full"}), controllerProjects.Delete)
	}
}
