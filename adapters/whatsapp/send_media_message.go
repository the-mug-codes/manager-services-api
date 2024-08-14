package whatsapp

import (
	"encoding/json"
)

type SendMediaMessage struct {
	Link    string  `json:"link" binding:"required"`
	Caption *string `json:"caption,omitempty"`
}

type whatsappSendMediaMessage struct {
	SendMessageRequest
	Audio    *SendMediaMessage `json:"audio,omitempty"`
	Document *SendMediaMessage `json:"document,omitempty"`
	Image    *SendMediaMessage `json:"image,omitempty"`
	Sticker  *SendMediaMessage `json:"sticker,omitempty"`
	Video    *SendMediaMessage `json:"video,omitempty"`
}

func (whatsapp *whatsapp) SendMediaMessage(to string, mediaType MediaType, mediaMessage SendMediaMessage) (err error) {
	messageBody := whatsappSendMediaMessage{
		SendMessageRequest: SendMessageRequest{
			MessagingProduct: "whatsapp",
			To:               to,
		},
	}
	switch mediaType {
	case MediaTypeAudio:
		messageBody.SendMessageRequest.Type = MessageTypeAudio
		messageBody.Audio = &mediaMessage
	case MediaTypeDocument:
		messageBody.SendMessageRequest.Type = MessageTypeDocument
		messageBody.Document = &mediaMessage
	case MediaTypeImage:
		messageBody.SendMessageRequest.Type = MessageTypeImage
		messageBody.Image = &mediaMessage
	case MediaTypeSticker:
		messageBody.SendMessageRequest.Type = MessageTypeSticker
		messageBody.Sticker = &mediaMessage
	case MediaTypeVideo:
		messageBody.SendMessageRequest.Type = MessageTypeVideo
		messageBody.Video = &mediaMessage
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
