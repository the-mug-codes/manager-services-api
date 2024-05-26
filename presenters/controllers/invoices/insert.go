package attachment

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	invoice "github.com/kodit-tecnologia/service-manager/use_cases/invoice"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type insertInvoice struct {
	ContractID    *uuid.UUID `json:"contract_id"`
	ContractTitle string     `json:"contract_title"`
	PersonID      uuid.UUID  `json:"person_id"`
	Items         []*struct {
		ProductID   uuid.UUID `json:"product_id" binding:"required"`
		PriceValue  float64   `json:"price_value" binding:"required"`
		Quantity    float64   `json:"quantity" binding:"required"`
		TotalAmount float64   `json:"total_amount" binding:"required"`
	} `json:"items"`
	AmountValue         float64 `json:"amount_value"`
	DiscountValue       float64 `json:"discount_value"`
	ExtraAmountValue    float64 `json:"extra_amount_value"`
	TotalAmountValue    float64 `json:"total_amount_value"`
	InstallmentNumber   *int    `json:"installment_number"`
	Comments            *string `json:"comments"`
	DiscountComments    *string `json:"discount_comments"`
	ExtraAmountComments *string `json:"extra_amount_comments"`
	WarningComments     *string `json:"warning_comments"`
	LateInvoices        *bool   `json:"late_invoices"`
}

func (insertInvoice) getDueDate() (dueDate time.Time) {
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

func (data insertInvoice) dataToInsert() (dataToInsert *entity.Invoice) {
	dataToInsert = &entity.Invoice{
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
		LateInvoices:        data.LateInvoices,
	}
	for _, item := range data.Items {
		dataToInsert.Items = append(dataToInsert.Items, &entity.InvoiceItem{
			ProductID:   item.ProductID,
			PriceValue:  item.PriceValue,
			Quantity:    item.Quantity,
			TotalAmount: item.TotalAmount,
		})
	}
	return dataToInsert
}

// @Summary		Create a Invoice
// @Description	Creates a new invoice in database.
// @Tags			Invoices
// @Accept			json
// @Produce		json
// @Param			payload	body		insertInvoice	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Invoice]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/invoices [post]
func Insert(context *gin.Context) {
	var requestBody *insertInvoice
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := invoice.Insert(repository.Invoice(context), *requestBody.dataToInsert())
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
