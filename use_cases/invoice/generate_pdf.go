package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func GeneratePDF(html entity.InvoiceHtml, pdf entity.InvoicePDF, invoice entity.Invoice) (file *[]byte, err error) {
	htmlString, err := html.Generate("document", invoice)
	if err != nil {
		return file, err
	}
	pdfDocumentPath, err := pdf.GenerateFile(invoice.ID.String(), *htmlString, invoice.Code, true, false, "A4", "Portrait")
	if err != nil {
		return file, err
	}
	return pdf.GenerateBinary(pdfDocumentPath, true)
}
