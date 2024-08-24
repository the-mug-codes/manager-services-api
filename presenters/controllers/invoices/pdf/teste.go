package attachment

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	html "github.com/the-mug-codes/adapters-service-api/html"
	pdf "github.com/the-mug-codes/adapters-service-api/pdf"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	pix "github.com/the-mug-codes/service-manager-api/adapters/pix"
)

type teste struct {
	PixCopyPaste *string `json:"pix_copy_paste,omitempty"`
	PixQrCode    *string `json:"pix_qr_code,omitempty"`
}

func Teste(context *gin.Context) {

	pixKey := "46415877000144"
	pixName := "Koditec Inova Simples"
	pixCity := "Curitiba"
	value := 89.9
	copyPaste, qrCode, err := pix.Start(pixKey, pixName, pixCity).CreatePix(&value, nil, nil)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert payment", err.Error())
		return
	}
	qrCode = qrCode[22:]
	invoice := &teste{}
	invoice.PixCopyPaste = &copyPaste
	invoice.PixQrCode = &qrCode
	html := html.Html[teste]("templates")
	pdf := pdf.Pdf("tmp")
	htmlString, err := html.Generate("teste", *invoice)
	if err != nil {
		panic(err)
	}
	pdfDocumentPath, _ := pdf.GenerateFile(uuid.NewString(), *htmlString, "INV-b52c674e-bcef", true, false, "A4", "Portrait")
	file, _ := pdf.GenerateBinary(pdfDocumentPath, true)
	context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", "INV-9ec39008-b9f7"))
	context.Data(200, "application/pdf", *file)
}
