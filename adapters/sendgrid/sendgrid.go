package sendgrid

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SendGridInterface[EmailAttachment any] interface {
	SendEmailMessage(to string, name string, subject string, text string, html *string, attachments *[]EmailAttachment) (err error)
}

type sendGrid[EmailAttachment any] struct {
	FromName   string
	SendEmail  string
	ReplyEmail string
}

func Connect[EmailAttachment any](fromName string, sendEmail string, replyEmail string) SendGridInterface[EmailAttachment] {
	return &sendGrid[EmailAttachment]{
		FromName:   fromName,
		SendEmail:  sendEmail,
		ReplyEmail: replyEmail,
	}
}

func (sendGrid *sendGrid[EmailAttachment]) apiRequest(method string, path string, requestBody *[]byte, queryParams *[]map[string]string) (responseBody []byte, err error) {
	body := &bytes.Reader{}
	if requestBody != nil {
		body = bytes.NewReader(*requestBody)
	}
	client := http.Client{}
	url := fmt.Sprintf("https://api.sendgrid.com%s", path)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return responseBody, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SEND_GRID_KEY")))
	if queryParams != nil {
		query := request.URL.Query()
		for _, paramMap := range *queryParams {
			for key, value := range paramMap {
				query.Add(key, value)
			}
		}
		request.URL.RawQuery = query.Encode()
	}
	response, err := client.Do(request)
	if err != nil {
		return responseBody, err
	}
	defer response.Body.Close()
	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return responseBody, err
	}
	return responseBody, err
}
