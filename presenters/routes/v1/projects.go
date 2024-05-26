package v1

import (
	"github.com/gin-gonic/gin"
	controllerProjects "github.com/kodit-tecnologia/service-manager/presenters/controllers/projects"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Projects(router *gin.RouterGroup) {
	projectsRoute := router.Group("projects")
	{
		projectsRoute.POST("", controllerProjects.Insert)
		projectsRoute.GET("", controllerProjects.ReadAll)
		projectsRoute.GET(":id", controllerProjects.Read)
		projectsRoute.PUT(":id", controllerProjects.Update)
		projectsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerProjects.Delete)
	}
}
