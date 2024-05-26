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
				messages.POST("", middleware.Protected(nil, nil), controllerTicketsMessages.Insert)
				messages.GET("", middleware.Protected(nil, nil), controllerTicketsMessages.ReadAll)
				messages.GET(":id/messages", middleware.Protected(nil, nil), controllerTicketsMessages.ReadAllByTicket)
				messages.GET(":id", middleware.Protected(nil, nil), controllerTicketsMessages.Read)
			}
			ticketsRoute.POST("", middleware.Protected(nil, nil), controllerTickets.Insert)
			ticketsRoute.GET("", middleware.Protected(nil, nil), controllerTickets.ReadAll)
			ticketsRoute.GET(":id", middleware.Protected(nil, nil), controllerTickets.Read)
			ticketsRoute.PUT(":id", middleware.Protected(nil, nil), controllerTickets.Update)
			ticketsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerTickets.Delete)
		}
	}
}
