package v1

import (
	"github.com/gin-gonic/gin"
	controllerChat "github.com/the-mug-codes/service-manager-api/presenters/controllers/chat"
)

func Chat(router *gin.RouterGroup) {
	chatRoute := router.Group("chat")
	{
		sessionsRoute := chatRoute.Group("sessions")
		{
			sessionsRoute.POST("", controllerChat.CreateSection)
			sessionsRoute.GET("join/:id", controllerChat.JoinSection)
			sessionsRoute.GET("", controllerChat.ReadSections)
			sessionsRoute.GET(":id/messages", controllerChat.ReadMessages)
			sessionsRoute.DELETE(":id", controllerChat.DeleteSection)
		}
	}
}
