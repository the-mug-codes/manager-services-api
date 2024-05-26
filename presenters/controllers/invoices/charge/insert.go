package attachment

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	messageBird "github.com/kodit-tecnologia/service-manager/adapters/messagebird"
	pix "github.com/kodit-tecnologia/service-manager/adapters/pix"
	sendGrid "github.com/kodit-tecnologia/service-manager/adapters/sendgrid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	invoice "github.com/kodit-tecnologia/service-manager/use_cases/invoice"
	html "github.com/the-mug-codes/adapters-service-api/html"
	pdf "github.com/the-mug-codes/adapters-service-api/pdf"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type insertCharge struct {
	ContractID          *uuid.UUID            `json:"contract_id"`
	ContractTitle       string                `json:"contract_title"`
	PersonID            uuid.UUID             `json:"person_id"`
	Items               []*entity.InvoiceItem `json:"items"`
	AmountValue         float64               `json:"amount_value"`
	DiscountValue       float64               `json:"discount_value"`
	ExtraAmountValue    float64               `json:"extra_amount_value"`
	TotalAmountValue    float64               `json:"total_amount_value"`
	InstallmentNumber   *int                  `json:"installment_number"`
	Comments            *string               `json:"comments"`
	DiscountComments    *string               `json:"discount_comments"`
	ExtraAmountComments *string               `json:"extra_amount_comments"`
	WarningComments     *string               `json:"warning_comments"`
	LateInvoices        *bool                 `json:"late_invoices"`
	Payment             struct {
		BankSlip *bool `json:"bank_slip"`
		Pix      *bool `json:"pix"`
	} `json:"payment"`
	Notification struct {
		Email    bool `json:"email"`
		WhatsApp bool `json:"whatsapp"`
	} `json:"notification"`
}

func (insertCharge) getDueDate() (dueDate time.Time) {
	now := time.Now()
	dueDate = time.Date(now.Year(), now.Month(), 5, 0, 0, 0, 0, now.Location())
	dueDate = dueDate.AddDate(0, 1, 0)
	if dueDate.Day() == 0 {
		dueDate = dueDate.AddDate(0, 0, 1)
	}
	if dueDate.Day() == 6 {
		dueDate = dueDate.AddDate(0, 0, 2)
	}
	return dueDate
}

func (data insertCharge) dataToInsert() (dataToInsert *entity.Invoice) {
	return &entity.Invoice{
		ContractID:          data.ContractID,
		PersonID:            data.PersonID,
		DueDate:             data.getDueDate(),
		AmountValue:         data.AmountValue,
		DiscountValue:       data.DiscountValue,
		ExtraAmountValue:    data.ExtraAmountValue,
		TotalAmountValue:    data.TotalAmountValue,
		InstallmentNumber:   data.InstallmentNumber,
		Comments:            data.Comments,
		ExtraAmountComments: data.ExtraAmountComments,
		DiscountComments:    data.DiscountComments,
		WarningComments:     data.WarningComments,
		Items:               data.Items,
		LateInvoices:        data.LateInvoices,
	}
}

// @Summary		Create a Invoice Charge
// @Description	Creates a new invoice charge by invoice id in database.
// @Tags			Invoices
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID	true	"ID"
// @Param			payload	body		insertCharge	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Invoice]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/invoices/charge/:id [post]
func Insert(context *gin.Context) {
	var requestBody *insertCharge
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	dataToInsert := requestBody.dataToInsert()
	if requestBody.Payment.Pix != nil && *requestBody.Payment.Pix {
		pixKey := "46415877000144"
		pixName := "Koditec Inova Simples"
		pixCity := "Curitiba"
		copyPaste, qrCode, err := pix.Start(pixKey, pixName, pixCity).CreatePix(&dataToInsert.AmountValue, &requestBody.ContractTitle, &dataToInsert.Code)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot insert invoice charge", err.Error())
			return
		}
		dataToInsert.PixCopyPaste = &copyPaste
		dataToInsert.PixQrCode = &qrCode
	}
	response, err := invoice.Insert(repository.Invoice(context), *dataToInsert)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert invoice charge", err.Error())
		return
	}
	if requestBody.Notification.WhatsApp {
		dataToInsert = response
		messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
		err := invoice.SendWhatsAppNotification(messageBirdConnection, *dataToInsert)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot insert invoice charge", err.Error())
			return
		}
	}
	if requestBody.Notification.Email {
		sendGridConnection := sendGrid.Connect[entity.EmailAttachment]("The Mug Codes", "notification@the.mug.codes", "contato@the.mug.codes")
		htmlConnection := html.Html[entity.Invoice]("templates/invoice")
		pdfConnection := pdf.Pdf("tmp")
		err := invoice.SendEmailNotification(sendGridConnection, htmlConnection, pdfConnection, *dataToInsert)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot send message", err.Error())
			return
		}
	}
	helper.SuccessResponseOne(context, 201, response)
}
