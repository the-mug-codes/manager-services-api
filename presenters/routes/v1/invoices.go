package v1

import (
	"github.com/gin-gonic/gin"
	controllerInvoices "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices"
	controllerInvoicesCharge "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/charge"
	controllerInvoicesNotification "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/notification"
	controllerInvoicesPayment "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/payment"
	controllerInvoicesPDF "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/pdf"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Invoices(router *gin.RouterGroup) {
	invoicesRoute := router.Group("invoices")
	{
		invoicesPDFRoute := invoicesRoute.Group("pdf")
		{
			invoicesPDFRoute.GET("teste", middleware.Protected(nil, nil), controllerInvoicesPDF.Teste)
			invoicesPDFRoute.GET(":id/:filename.pdf", middleware.Protected(nil, nil), controllerInvoicesPDF.Read)
		}
		invoicesPaymentRoute := invoicesRoute.Group("payments")
		{
			invoicesPaymentRoute.POST(":id", middleware.Protected(nil, nil), controllerInvoicesPayment.Insert)
		}
		invoicesNotificationRoute := invoicesRoute.Group("notifications")
		{
			invoicesNotificationRoute.POST(":id", middleware.Protected(nil, nil), controllerInvoicesNotification.Insert)
		}
		invoicesChargeRoute := invoicesRoute.Group("charge")
		{
			invoicesChargeRoute.POST("", middleware.Protected(nil, nil), controllerInvoicesCharge.Insert)
		}
		invoicesRoute.POST("", middleware.Protected(nil, nil), controllerInvoices.Insert)
		invoicesRoute.GET(":id", middleware.Protected(nil, nil), controllerInvoices.Read)
		invoicesRoute.GET("", middleware.Protected(nil, nil), controllerInvoices.ReadAll)
	}
}
