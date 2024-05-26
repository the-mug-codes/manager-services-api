package attachment

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	invoice "github.com/kodit-tecnologia/service-manager/use_cases/invoice"
	html "github.com/the-mug-codes/adapters-service-api/html"
	pdf "github.com/the-mug-codes/adapters-service-api/pdf"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

// @Summary		Show a Invoice PDF File
// @Description	Get a invoice pdf file by idin database.
// @Tags			Invoices
// @Produce		application/pdf
// @Param			id	path		uuid.UUID	true	"ID"
// @Success 200 {string} application/pdf
// @Failure		400	{object}	helper.Error
// @Failure		401	{object}	helper.Error
// @Failure		404	{object}	helper.Error
// @Router			/invoices/pdf/{id} [get]
func Read(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot read", "id not provided")
		return
	}
	invoiceID, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	invoiceData, err := invoice.Read(repository.Invoice(context), invoiceID)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot read", err.Error())
		return
	}
	htmlConnection := html.Html[entity.Invoice]("templates/invoice")
	pdfConnection := pdf.Pdf("tmp")
	file, err := invoice.GeneratePDF(htmlConnection, pdfConnection, *invoiceData)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert", err.Error())
		return
	}
	context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", invoiceData.Code))
	context.Data(200, "application/pdf", *file)
}
