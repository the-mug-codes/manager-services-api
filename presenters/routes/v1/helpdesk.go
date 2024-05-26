package v1

import (
	"github.com/gin-gonic/gin"
	controllerTickets "github.com/kodit-tecnologia/service-manager/presenters/controllers/helpdesk/tickets"
	controllerTicketsMessages "github.com/kodit-tecnologia/service-manager/presenters/controllers/helpdesk/tickets/messages"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func HelpDesk(router *gin.RouterGroup) {
	helpdeskRoute := router.Group("helpdesk")
	{
		ticketsRoute := helpdeskRoute.Group("tickets")
		{
			messages := helpdeskRoute.Group("messages")
			{
				messages.POST("", controllerTicketsMessages.Insert)
				messages.GET("", controllerTicketsMessages.ReadAll)
				messages.GET(":id/messages", controllerTicketsMessages.ReadAllByTicket)
				messages.GET(":id", controllerTicketsMessages.Read)
			}
			ticketsRoute.POST("", controllerTickets.Insert)
			ticketsRoute.GET("", controllerTickets.ReadAll)
			ticketsRoute.GET(":id", controllerTickets.Read)
			ticketsRoute.PUT(":id", controllerTickets.Update)
			ticketsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerTickets.Delete)
		}
	}
}
