package whatsapp

import "encoding/json"

type interactiveType string

const (
	InteractiveTypeButton interactiveType = "button"
	InteractiveTypeList   interactiveType = "list"
)

type action struct {
	Button  *string `json:"button,omitempty"`
	Buttons *[]struct {
		Type  string `json:"type" binding:"required"`
		Reply *struct {
			ID    string `json:"id" binding:"required"`
			Title string `json:"title" binding:"required"`
		} `json:"reply,omitempty"`
	} `json:"buttons,omitempty"`
	Sections *[]struct {
		Title string `json:"title" binding:"required"`
		Rows  []struct {
			ID          string  `json:"id" binding:"required"`
			Title       string  `json:"title" binding:"required"`
			Description *string `json:"description,omitempty"`
		} `json:"rows" binding:"required"`
	} `json:"sections,omitempty"`
}

type SendInteractiveMessage struct {
	Type   interactiveType `json:"type" binding:"required"`
	Action action          `json:"action" binding:"required"`
	Body   *struct {
		Text string `json:"text" binding:"required"`
	} `json:"body,omitempty"`
	Header *struct {
		Text string `json:"text" binding:"required"`
	} `json:"header,omitempty"`
	Footer *struct {
		Text string `json:"text" binding:"required"`
	} `json:"footer,omitempty"`
}

type whatsappSendInteractiveMessage struct {
	SendMessageRequest
	Interactive SendInteractiveMessage `json:"interactive"`
}

func (whatsapp *whatsapp) SendInteractiveMessage(to string, interactiveMessage SendInteractiveMessage) (err error) {
	messageBody := whatsappSendInteractiveMessage{
		SendMessageRequest: SendMessageRequest{
			MessagingProduct: "whatsapp",
			Type:             MessageTypeInteractive,
			To:               to,
		},
		Interactive: interactiveMessage,
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
