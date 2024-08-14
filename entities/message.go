package entities

type WhatsappMessageType string

const (
	WhatsappMessageTypeText        WhatsappMessageType = "text"
	WhatsappMessageTypeInteractive WhatsappMessageType = "interactive"
	WhatsappMessageTypeTemplate    WhatsappMessageType = "template"
	WhatsappMessageTypeAudio       WhatsappMessageType = "audio"
	WhatsappMessageTypeDocument    WhatsappMessageType = "document"
	WhatsappMessageTypeImage       WhatsappMessageType = "image"
	WhatsappMessageTypeSticker     WhatsappMessageType = "sticker"
	WhatsappMessageTypeVideo       WhatsappMessageType = "video"
)

type WhatsappMessage struct {
	MessageType     WhatsappMessageType `json:"message_type" binding:"required"`
	Template        *string             `json:"template" binding:"omitempty"`
	Language        *string             `json:"language" binding:"omitempty"`
	HeaderVariables *[]string           `json:"header_variables" binding:"omitempty"`
	BodyVariables   *[]string           `json:"body_variables" binding:"omitempty"`
	ButtonVariables *[]string           `json:"button_variables" binding:"omitempty"`
	Header          *string             `json:"header" binding:"omitempty"`
	Body            *string             `json:"body" binding:"omitempty"`
	Footer          *string             `json:"footer" binding:"omitempty"`
	Buttons         *[]string           `json:"buttons" binding:"omitempty"`
	URL             *string             `json:"url" binding:"omitempty"`
}

type EmailAttachment struct {
	Content     string `json:"content"`
	Filename    string `json:"filename"`
	Type        string `json:"type"`
	Disposition string `json:"disposition"`
}

type NewMessageCreated struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type MessageContent struct {
	Text  *string `json:"text"`
	Image *struct {
		Url     string `json:"url"`
		Caption string `json:"caption"`
	} `json:"image"`
	Video *struct {
		Url     string `json:"url"`
		Caption string `json:"caption"`
	} `json:"video"`
	Audio *struct {
		Url     string `json:"url"`
		Caption string `json:"caption"`
	} `json:"audio"`
	File *struct {
		Url     string `json:"url"`
		Caption string `json:"caption"`
	} `json:"file"`
	Hsm *struct {
		Namespace    string `json:"namespace"`
		TemplateName string `json:"templateName"`
		Language     struct {
			Code string `json:"code"`
		} `json:"language"`
		Params *[]struct {
			Default string `json:"default"`
		} `json:"params"`
		Components *[]struct {
			Type       string `json:"type"`
			SubType    string `json:"sub_type"`
			Parameters *[]struct {
				Type     string  `json:"type"`
				Text     *string `json:"text"`
				Document *struct {
					Url string `json:"url"`
				} `json:"document"`
				Image *struct {
					Url string `json:"url"`
				} `json:"image"`
				Video *struct {
					Url string `json:"url"`
				} `json:"video"`
			} `json:"parameters"`
		} `json:"components"`
	} `json:"hsm"`
	Interactive *struct {
		Type string `json:"type"`
		Body struct {
			Text string `json:"text"`
		} `json:"body"`
		Footer *struct {
			Text string `json:"text"`
		} `json:"footer"`
		Reply *struct {
			Text string `json:"text"`
		} `json:"reply"`
	} `json:"interactive"`
	Email *struct {
		To []struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"to"`
		From struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"from"`
		Subject string `json:"subject"`
		Content struct {
			Html string `json:"html"`
			Text string `json:"text"`
		} `json:"content"`
		Attachments *[]struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Url  string `json:"url"`
		} `json:"attachments"`
	} `json:"email"`
}
