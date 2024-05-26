package sendgrid

import (
	"encoding/json"
)

type emailRecipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type emailContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type emailMessage[EmailAttachment any] struct {
	Personalizations []struct {
		To      []emailRecipient `json:"to"`
		Subject string           `json:"subject"`
	} `json:"personalizations"`
	Content     []emailContent     `json:"content"`
	From        emailRecipient     `json:"from"`
	ReplyTo     emailRecipient     `json:"reply_to"`
	Attachments *[]EmailAttachment `json:"attachments"`
}

func (sendGrid *sendGrid[EmailAttachment]) SendEmailMessage(to string, name string, subject string, text string, html *string, attachments *[]EmailAttachment) (err error) {
	content := []emailContent{
		{
			Type:  "text/plain",
			Value: text,
		},
	}
	if html != nil {
		content = append(content, emailContent{
			Type:  "text/html",
			Value: *html,
		})
	}
	email := emailMessage[EmailAttachment]{
		Personalizations: []struct {
			To      []emailRecipient `json:"to"`
			Subject string           `json:"subject"`
		}{
			{
				To: []emailRecipient{
					{
						Email: to,
					},
				},
				Subject: subject,
			},
		},
		Content:     content,
		Attachments: attachments,
		From: emailRecipient{
			Email: sendGrid.SendEmail,
			Name:  sendGrid.FromName,
		},
		ReplyTo: emailRecipient{
			Email: sendGrid.ReplyEmail,
			Name:  sendGrid.FromName,
		},
	}
	requestBody, err := json.Marshal(email)
	if err != nil {
		return err
	}
	_, err = sendGrid.apiRequest("POST", "/v3/mail/send", &requestBody, nil)
	if err != nil {
		return err
	}
	return err
}
