package whatsapp

import (
	"encoding/json"
)

type componentsType string

const (
	ComponentTypeHeader componentsType = "header"
	ComponentTypeBody   componentsType = "body"
	ComponentTypeButton componentsType = "button"
)

type componentsSubType string

const (
	ComponentSubTypeQuickReply componentsSubType = "quick_reply"
	ComponentSubTypeURL        componentsSubType = "url"
	ComponentSubTypeCatalog    componentsSubType = "catalog"
)

type parametersType string

const (
	ParameterTypeButton   parametersType = "button"
	ParameterTypeDocument parametersType = "document"
	ParameterTypeImage    parametersType = "image"
	ParameterTypeText     parametersType = "text"
	ParameterTypeVideo    parametersType = "video"
)

type language struct {
	Code string `json:"code"`
}

type parameters struct {
	Type     parametersType    `json:"type"`
	Text     *parametersType   `json:"text,omitempty"`
	Audio    *SendMediaMessage `json:"audio,omitempty"`
	Document *SendMediaMessage `json:"document,omitempty"`
	Image    *SendMediaMessage `json:"image,omitempty"`
	Sticker  *SendMediaMessage `json:"sticker,omitempty"`
	Video    *SendMediaMessage `json:"video,omitempty"`
	Button   *struct {
		Type    string  `json:"type"`
		Payload *string `json:"payload"`
		Text    *string `json:"text"`
	} `json:"button,omitempty"`
}

type components struct {
	Type       componentsType     `json:"type"`
	Parameters *[]parameters      `json:"parameters"`
	SubType    *componentsSubType `json:"sub_type,omitempty"`
	Index      *int               `json:"index,omitempty"`
}

type SendTemplateMessage struct {
	Language   language      `json:"language" binding:"required"`
	Name       string        `json:"name" binding:"required"`
	Components *[]components `json:"components,omitempty"`
}

type whatsappSendTemplateMessage struct {
	SendMessageRequest
	Template SendTemplateMessage `json:"template" binding:"required"`
}

func (whatsapp *whatsapp) SendTemplateMessage(to string, templateMessage SendTemplateMessage) (err error) {
	messageBody := whatsappSendTemplateMessage{
		SendMessageRequest: SendMessageRequest{
			MessagingProduct: "whatsapp",
			Type:             MessageTypeTemplate,
			To:               to,
		},
		Template: templateMessage,
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
