package v1

import (
	"github.com/gin-gonic/gin"
	controllerChat "github.com/the-mug-codes/service-manager-api/presenters/controllers/chat"
)

func Chat(router *gin.RouterGroup) {
	chatRoute := router.Group("chat")
	{
		chatRoute.POST("create", controllerChat.CreateSection)
		chatRoute.GET("sessions", controllerChat.ReadSections)
		chatRoute.GET("join/:id", controllerChat.JoinSection)
		chatRoute.DELETE(":id", controllerChat.DeleteSection)
	}
}
