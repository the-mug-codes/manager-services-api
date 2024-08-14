package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerTickets "github.com/the-mug-codes/service-manager-api/presenters/controllers/helpdesk/tickets"
	controllerTicketsMessages "github.com/the-mug-codes/service-manager-api/presenters/controllers/helpdesk/tickets/messages"
)

func HelpDesk(router *gin.RouterGroup) {
	helpdeskRoute := router.Group("helpdesk")
	{
		ticketsRoute := helpdeskRoute.Group("tickets")
		{
			messages := helpdeskRoute.Group("messages")
			{
				messages.POST("", middleware.Protected(&[]string{"admin:full", "user:full"}), controllerTicketsMessages.Insert)
				messages.GET("", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerTicketsMessages.ReadAll)
				messages.GET(":id/messages", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerTicketsMessages.ReadAllByTicket)
				messages.GET(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerTicketsMessages.Read)
			}
			ticketsRoute.POST("", middleware.Protected(&[]string{"admin:full", "user:full"}), controllerTickets.Insert)
			ticketsRoute.GET("", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerTickets.ReadAll)
			ticketsRoute.GET(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerTickets.Read)
			ticketsRoute.PUT(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerTickets.Update)
			ticketsRoute.DELETE(":id", middleware.Protected(&[]string{"admin:full"}), controllerTickets.Delete)
		}
	}
}
