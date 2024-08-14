package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	html "github.com/the-mug-codes/adapters-service-api/html"
	pdf "github.com/the-mug-codes/adapters-service-api/pdf"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	messageBird "github.com/the-mug-codes/service-manager-api/adapters/messagebird"
	sendGrid "github.com/the-mug-codes/service-manager-api/adapters/sendgrid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	invoice "github.com/the-mug-codes/service-manager-api/use_cases/invoice"
)

type insertNotification struct {
	Email    bool `json:"email"`
	WhatsApp bool `json:"whatsapp"`
}

// @Summary		Create a Invoice Notification
// @Description	Creates a new invoice notification by invoice id in database.
// @Tags			Invoices
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID	true	"ID"
// @Param			payload	body		insertNotification	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Invoice]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/invoices/notifications/:id [post]
func Insert(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot bind data", "id not provided")
		return
	}
	invoiceID, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	var request *insertNotification
	err = context.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	dataToInsert, err := invoice.Read(repository.Invoice(context), invoiceID)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	if request.WhatsApp {
		messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
		err := invoice.SendWhatsAppNotification(messageBirdConnection, *dataToInsert)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot send message", err.Error())
			return
		}
	}
	if request.Email {
		sendGridConnection := sendGrid.Connect[entity.EmailAttachment]("The Mug Codes", "notification@the.mug.codes", "contato@the.mug.codes")
		htmlConnection := html.Html[entity.Invoice]("templates/invoice")
		pdfConnection := pdf.Pdf("tmp")
		err := invoice.SendEmailNotification(sendGridConnection, htmlConnection, pdfConnection, *dataToInsert)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot send message", err.Error())
			return
		}
	}
	helper.SuccessResponseNone(context, 200)
}
