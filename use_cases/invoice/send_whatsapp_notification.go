package people

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func generateBase64File(file []byte) (base64File string, err error) {
	base64String := &strings.Builder{}
	_, err = base64String.WriteString("data:application/pdf;base64,")
	if err != nil {
		return base64File, err
	}
	_, err = base64String.WriteString(base64.StdEncoding.EncodeToString(file))
	if err != nil {
		return base64File, err
	}
	base64File = base64String.String()
	return base64File, err
}

func SendWhatsAppNotification(whatsapp entity.InvoiceWhatsApp, invoice entity.Invoice) (err error) {
	name := invoice.Person.GetFistName()
	totalAmountValue := invoice.ParseMoney(invoice.TotalAmountValue)
	dueDate := invoice.ParseDate(invoice.DueDate)
	fileLink := fmt.Sprintf("%s/v1/invoices/pdf/%v/%s.pdf", os.Getenv("HOST"), invoice.ID, invoice.Code)
	messageContent := entity.MessageContent{
		Hsm: &struct {
			Namespace    string "json:\"namespace\""
			TemplateName string "json:\"templateName\""
			Language     struct {
				Code string "json:\"code\""
			} "json:\"language\""
			Params *[]struct {
				Default string "json:\"default\""
			} "json:\"params\""
			Components *[]struct {
				Type       string "json:\"type\""
				SubType    string "json:\"sub_type\""
				Parameters *[]struct {
					Type     string  "json:\"type\""
					Text     *string "json:\"text\""
					Document *struct {
						Url string "json:\"url\""
					} "json:\"document\""
					Image *struct {
						Url string "json:\"url\""
					} "json:\"image\""
					Video *struct {
						Url string "json:\"url\""
					} "json:\"video\""
				} "json:\"parameters\""
			} "json:\"components\""
		}{
			Language: struct {
				Code string "json:\"code\""
			}{
				Code: "pt_BR",
			},
			Namespace:    "0bac3efc_3e40_4fa5_8791_20c83633d938",
			TemplateName: "send_invoice",
			Components: &[]struct {
				Type       string "json:\"type\""
				SubType    string "json:\"sub_type\""
				Parameters *[]struct {
					Type     string  "json:\"type\""
					Text     *string "json:\"text\""
					Document *struct {
						Url string "json:\"url\""
					} "json:\"document\""
					Image *struct {
						Url string "json:\"url\""
					} "json:\"image\""
					Video *struct {
						Url string "json:\"url\""
					} "json:\"video\""
				} "json:\"parameters\""
			}{
				{
					Type: "body",
					Parameters: &[]struct {
						Type     string  "json:\"type\""
						Text     *string "json:\"text\""
						Document *struct {
							Url string "json:\"url\""
						} "json:\"document\""
						Image *struct {
							Url string "json:\"url\""
						} "json:\"image\""
						Video *struct {
							Url string "json:\"url\""
						} "json:\"video\""
					}{
						{
							Type: "text",
							Text: &name,
						},
						{
							Type: "text",
							Text: &totalAmountValue,
						},
						{
							Type: "text",
							Text: &dueDate,
						},
					},
				},
				{
					Type: "header",
					Parameters: &[]struct {
						Type     string  "json:\"type\""
						Text     *string "json:\"text\""
						Document *struct {
							Url string "json:\"url\""
						} "json:\"document\""
						Image *struct {
							Url string "json:\"url\""
						} "json:\"image\""
						Video *struct {
							Url string "json:\"url\""
						} "json:\"video\""
					}{
						{
							Type: "document",
							Document: &struct {
								Url string "json:\"url\""
							}{Url: fileLink},
						},
					},
				},
			},
		},
	}
	_, err = whatsapp.SendMessage(invoice.Person.GetMainPhone(true).FullPhone, "fab03183-b828-4088-9f11-226137ebf9dd", "hsm", messageContent, nil)
	if err != nil {
		return err
	}
	return err
}
