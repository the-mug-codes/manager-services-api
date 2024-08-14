package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerWhatsapp "github.com/the-mug-codes/service-manager-api/presenters/controllers/whatsapp"
)

func Whatsapp(router *gin.RouterGroup) {
	whatsappRoute := router.Group("whatsapp")
	{
		webhooks := whatsappRoute.Group("webhooks")
		{
			webhooks.GET("", controllerWhatsapp.Subscribe)
			webhooks.POST("", controllerWhatsapp.MessageReceived)
		}
		messages := whatsappRoute.Group("messages")
		{
			messages.GET("media/:id", middleware.Protected(&[]string{"admin:full"}), controllerWhatsapp.ReadMedia)
			messages.POST(":phone/text", middleware.Protected(&[]string{"admin:full"}), controllerWhatsapp.SendTextMessage)
			messages.POST(":phone/media", middleware.Protected(&[]string{"admin:full"}), controllerWhatsapp.SendMediaMessage)
			messages.POST(":phone/interactive", middleware.Protected(&[]string{"admin:full"}), controllerWhatsapp.SendInteractiveMessage)
			messages.POST(":phone/template", middleware.Protected(&[]string{"admin:full"}), controllerWhatsapp.SendTemplateMessage)
		}

	}
}
