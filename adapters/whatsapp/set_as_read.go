package whatsapp

import "encoding/json"

type messageStatusRequest struct {
	MessagingProduct string `json:"messaging_product"`
	Status           string `json:"status"`
	MessageID        string `json:"message_id"`
}

func (whatsapp *whatsapp) SetAsRead(messageID string) (err error) {
	messageStatus := messageStatusRequest{
		MessagingProduct: "whatsapp",
		Status:           "read",
		MessageID:        messageID,
	}
	requestBody, err := json.Marshal(messageStatus)
	if err != nil {
		return err
	}
	_, err = whatsapp.apiRequest("POST", "messages", &requestBody)
	if err != nil {
		return err
	}
	return err
}
