package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/the-mug-codes/service-manager-api/presenters/routes/v1"
)

func Routes(router *gin.Engine) {
	main := router.Group("v1")
	{
		v1.Contracts(main)
		v1.HelpDesk(main)
		v1.Invoices(main)
		v1.Labels(main)
		v1.MessageBird(main)
		v1.People(main)
		v1.Projects(main)
		v1.Products(main)
		v1.Security(main)
		v1.Whatsapp(main)
		v1.Chat(main)
	}
}
