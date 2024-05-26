package people

import (
	"fmt"

	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func SendEmailNotification(sendGrid entity.InvoiceEmail, html entity.InvoiceHtml, pdf entity.InvoicePDF, invoice entity.Invoice) (err error) {
	htmlString, err := html.Generate("invoice", invoice)
	if err != nil {
		return err
	}
	emailHtmlString, err := html.Generate("invoice_email", invoice)
	if err != nil {
		return err
	}
	pdfDocumentPath, err := pdf.GenerateFile(invoice.ID.String(), *htmlString, invoice.Code, true, false, "A4", "Portrait")
	if err != nil {
		return err
	}
	fileContent, err := pdf.GenerateBase64(pdfDocumentPath, true, false)
	if err != nil {
		return err
	}
	subject := fmt.Sprintf("[mug] Sua fatura com vencimento em %s já está fechada!", invoice.ParseDate(invoice.DueDate))
	text := fmt.Sprintf("Olá %s,\n\nSua fatura no valor de %s foi fechada e está pronta para pagamento até o vencimento em %s. Por favor, consulte o documento PDF em anexo para mais detalhes. Este e-mail é apenas uma notificação, não requer resposta.\n\nAtenciosamente,\nThe Mug Codes\ncontato@the.mug.codes\n+55 (41) 3891-0372\n\n(PT) Esta mensagem pode conter informação confidencial ou privilegiada, sendo seu sigilo protegido por lei. Se você não for o destinatário ou a pessoa autorizada a receber esta mensagem, não pode usar, copiar ou divulgar as informações nela contidas ou tomar qualquer ação baseada nessas informações. Se você recebeu esta mensagem por engano, por favor, avise imediatamente ao remetente, respondendo o e-mail e em seguida apague-a. Agradecemos sua cooperação.\n\n(EN) This message may contain confidential or privileged information and its confidentiality is protected by law. If you are not the addressed or authorized person to receive this message, you must not use, copy, disclose or take any action based on it or any information herein. If you have received this message by mistake, please advise the sender immediately by replying the e-mail and then deleting it. Thank you for your cooperation.", invoice.Person.GetFistName(), invoice.ParseMoney(invoice.TotalAmountValue), invoice.ParseDate(invoice.DueDate))
	attachments := &[]entity.EmailAttachment{
		{
			Content:     *fileContent,
			Filename:    fmt.Sprintf("%s.pdf", invoice.Code),
			Type:        "application/pdf",
			Disposition: "attachment",
		},
	}
	err = sendGrid.SendEmailMessage(invoice.Person.GetMainEmail().Email, invoice.Person.Name, subject, text, emailHtmlString, attachments)
	if err != nil {
		return err
	}
	return err
}
