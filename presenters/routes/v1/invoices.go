package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerInvoices "github.com/the-mug-codes/service-manager-api/presenters/controllers/invoices"
	controllerInvoicesCharge "github.com/the-mug-codes/service-manager-api/presenters/controllers/invoices/charge"
	controllerInvoicesNotification "github.com/the-mug-codes/service-manager-api/presenters/controllers/invoices/notification"
	controllerInvoicesPayment "github.com/the-mug-codes/service-manager-api/presenters/controllers/invoices/payment"
	controllerInvoicesPDF "github.com/the-mug-codes/service-manager-api/presenters/controllers/invoices/pdf"
)

func Invoices(router *gin.RouterGroup) {
	invoicesRoute := router.Group("invoices")
	{
		invoicesPDFRoute := invoicesRoute.Group("pdf")
		{
			invoicesPDFRoute.GET("teste", controllerInvoicesPDF.Teste)
			invoicesPDFRoute.GET(":id/:filename.pdf", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerInvoicesPDF.Read)
		}
		invoicesPaymentRoute := invoicesRoute.Group("payments")
		{
			invoicesPaymentRoute.POST(":id", middleware.Protected(&[]string{"admin:full"}), controllerInvoicesPayment.Insert)
		}
		invoicesNotificationRoute := invoicesRoute.Group("notifications")
		{
			invoicesNotificationRoute.POST(":id", middleware.Protected(&[]string{"admin:full"}), controllerInvoicesNotification.Insert)
		}
		invoicesChargeRoute := invoicesRoute.Group("charge")
		{
			invoicesChargeRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerInvoicesCharge.Insert)
		}
		invoicesRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerInvoices.Insert)
		invoicesRoute.GET(":id", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerInvoices.Read)
		invoicesRoute.GET("", middleware.Protected(&[]string{"admin:full", "user:self"}), controllerInvoices.ReadAll)
	}
}
