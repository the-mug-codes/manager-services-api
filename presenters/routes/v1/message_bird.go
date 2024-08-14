package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerMessageBird "github.com/the-mug-codes/service-manager-api/presenters/controllers/message_bird"
)

func MessageBird(router *gin.RouterGroup) {
	messageBirdRoute := router.Group("message-bird")
	{
		conversations := messageBirdRoute.Group("conversations")
		{
			conversations.POST("messages", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.SendMessage)
			conversations.GET("", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetAllConversations)
			conversations.GET(":id/messages", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetConversationMessages)
			conversations.GET(":id", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetConversation)
		}
		contacts := messageBirdRoute.Group("contacts")
		{
			contacts.GET("", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetAllContacts)
			contacts.GET(":id", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetContact)
		}
		calls := messageBirdRoute.Group("calls")
		{
			calls.GET("", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetAllCalls)
			calls.GET(":id/recordings", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetCallRecording)
			calls.GET(":id", middleware.Protected(&[]string{"admin:full"}), controllerMessageBird.GetCall)
		}
	}
}
