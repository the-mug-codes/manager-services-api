package v1

import (
	"github.com/gin-gonic/gin"
	controllerInvoices "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices"
	controllerInvoicesCharge "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/charge"
	controllerInvoicesNotification "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/notification"
	controllerInvoicesPayment "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/payment"
	controllerInvoicesPDF "github.com/kodit-tecnologia/service-manager/presenters/controllers/invoices/pdf"
)

func Invoices(router *gin.RouterGroup) {
	invoicesRoute := router.Group("invoices")
	{
		invoicesPDFRoute := invoicesRoute.Group("pdf")
		{
			invoicesPDFRoute.GET("teste", controllerInvoicesPDF.Teste)
			invoicesPDFRoute.GET(":id/:filename.pdf", controllerInvoicesPDF.Read)
		}
		invoicesPaymentRoute := invoicesRoute.Group("payments")
		{
			invoicesPaymentRoute.POST(":id", controllerInvoicesPayment.Insert)
		}
		invoicesNotificationRoute := invoicesRoute.Group("notifications")
		{
			invoicesNotificationRoute.POST(":id", controllerInvoicesNotification.Insert)
		}
		invoicesChargeRoute := invoicesRoute.Group("charge")
		{
			invoicesChargeRoute.POST("", controllerInvoicesCharge.Insert)
		}
		invoicesRoute.POST("", controllerInvoices.Insert)
		invoicesRoute.GET(":id", controllerInvoices.Read)
		invoicesRoute.GET("", controllerInvoices.ReadAll)
	}
}
