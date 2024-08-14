package whatsapp

import (
	"encoding/json"
)

type SendTextMessage struct {
	Body       string `json:"body" binding:"required"`
	PreviewUrl bool   `json:"preview_url,omitempty"`
}

type whatsappSendTextMessage struct {
	SendMessageRequest
	Text SendTextMessage `json:"text" binding:"required"`
}

func (whatsapp *whatsapp) SendTextMessage(to string, textMessage SendTextMessage) (err error) {
	messageBody := whatsappSendTextMessage{
		SendMessageRequest: SendMessageRequest{
			MessagingProduct: "whatsapp",
			Type:             MessageTypeText,
			To:               to,
		},
		Text: textMessage,
	}
	requestBody, err := json.Marshal(messageBody)
	if err != nil {
		return err
	}
	_, err = whatsapp.apiRequest("POST", "messages", &requestBody)
	if err != nil {
		return err
	}
	return err
}
