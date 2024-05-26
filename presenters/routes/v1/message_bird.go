package v1

import (
	"github.com/gin-gonic/gin"
	controllerMessageBird "github.com/kodit-tecnologia/service-manager/presenters/controllers/message_bird"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func MessageBird(router *gin.RouterGroup) {
	messageBirdRoute := router.Group("message-bird")
	{
		conversations := messageBirdRoute.Group("conversations")
		{
			conversations.POST("messages", middleware.Protected(nil, nil), controllerMessageBird.SendMessage)
			conversations.GET("", middleware.Protected(nil, nil), controllerMessageBird.GetAllConversations)
			conversations.GET(":id/messages", middleware.Protected(nil, nil), controllerMessageBird.GetConversationMessages)
			conversations.GET(":id", middleware.Protected(nil, nil), controllerMessageBird.GetConversation)
		}
		contacts := messageBirdRoute.Group("contacts")
		{
			contacts.GET("", middleware.Protected(nil, nil), controllerMessageBird.GetAllContacts)
			contacts.GET(":id", middleware.Protected(nil, nil), controllerMessageBird.GetContact)
		}
		calls := messageBirdRoute.Group("calls")
		{
			calls.GET("", middleware.Protected(nil, nil), controllerMessageBird.GetAllCalls)
			calls.GET(":id/recordings", middleware.Protected(nil, nil), controllerMessageBird.GetCallRecording)
			calls.GET(":id", middleware.Protected(nil, nil), controllerMessageBird.GetCall)
		}
	}
}
