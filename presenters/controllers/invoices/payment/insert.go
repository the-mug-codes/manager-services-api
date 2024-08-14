package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	pix "github.com/the-mug-codes/service-manager-api/adapters/pix"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	invoice "github.com/the-mug-codes/service-manager-api/use_cases/invoice"
)

type insertPayment struct {
	ContractTitle string `json:"project_name"`
	Payment       struct {
		BankSlip *bool `json:"bank_slip"`
		Pix      *bool `json:"pix"`
	} `json:"payment"`
}

// @Summary		Create a Invoice Payment Method
// @Description	Creates a new invoice payment method by invoice id in database.
// @Tags			Invoices
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID	true	"ID"
// @Param			payload	body		insertPayment	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Invoice]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/invoices/payments/:id [post]
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
	var request *insertPayment
	err = context.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	dataToInsert, err := invoice.Read(repository.Invoice(context), invoiceID)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert payment", err.Error())
		return
	}
	if request.Payment.Pix != nil && *request.Payment.Pix {
		pixKey := "46415877000144"
		pixName := "Koditec Inova Simples"
		pixCity := "Curitiba"
		copyPaste, qrCode, err := pix.Start(pixKey, pixName, pixCity).CreatePix(&dataToInsert.AmountValue, &request.ContractTitle, &dataToInsert.Code)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot insert payment", err.Error())
			return
		}
		dataToInsert.PixCopyPaste = &copyPaste
		dataToInsert.PixQrCode = &qrCode
	}
	response, err := invoice.Update(repository.Invoice(context), *dataToInsert)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert payment", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
