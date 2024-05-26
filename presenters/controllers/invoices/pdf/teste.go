package attachment

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pix "github.com/kodit-tecnologia/service-manager/adapters/pix"
	html "github.com/the-mug-codes/adapters-service-api/html"
	pdf "github.com/the-mug-codes/adapters-service-api/pdf"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type teste struct {
	PixCopyPaste *string `json:"pix_copy_paste,omitempty"`
	PixQrCode    *string `json:"pix_qr_code,omitempty"`
}

func Teste(context *gin.Context) {

	pixKey := "46415877000144"
	pixName := "Koditec Inova Simples"
	pixCity := "Curitiba"
	value := 269.7
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
	htmlString, _ := html.Generate("teste", *invoice)
	pdfDocumentPath, _ := pdf.GenerateFile(uuid.NewString(), *htmlString, "INV-468E-0506", true, false, "A4", "Portrait")
	file, _ := pdf.GenerateBinary(pdfDocumentPath, true)
	context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", "INV-468E-0506"))
	context.Data(200, "application/pdf", *file)
}
