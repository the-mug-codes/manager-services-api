package v1

import (
	"github.com/gin-gonic/gin"
	controllerProjects "github.com/kodit-tecnologia/service-manager/presenters/controllers/projects"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Projects(router *gin.RouterGroup) {
	projectsRoute := router.Group("projects")
	{
		projectsRoute.POST("", middleware.Protected(nil, nil), controllerProjects.Insert)
		projectsRoute.GET("", middleware.Protected(nil, nil), controllerProjects.ReadAll)
		projectsRoute.GET(":id", middleware.Protected(nil, nil), controllerProjects.Read)
		projectsRoute.PUT(":id", middleware.Protected(nil, nil), controllerProjects.Update)
		projectsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerProjects.Delete)
	}
}
