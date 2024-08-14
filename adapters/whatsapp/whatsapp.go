package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type MediaDownload struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	FileSize int    `json:"file_size"`
}

type MediaType string

const (
	MediaTypeAudio    MediaType = "audio"
	MediaTypeDocument MediaType = "document"
	MediaTypeImage    MediaType = "image"
	MediaTypeSticker  MediaType = "sticker"
	MediaTypeVideo    MediaType = "video"
)

type MessageType string

const (
	MessageTypeText        MessageType = "text"
	MessageTypeInteractive MessageType = "interactive"
	MessageTypeTemplate    MessageType = "template"
	MessageTypeAudio       MessageType = "audio"
	MessageTypeDocument    MessageType = "document"
	MessageTypeImage       MessageType = "image"
	MessageTypeSticker     MessageType = "sticker"
	MessageTypeVideo       MessageType = "video"
	MessageTypeUnknown     MessageType = "unknown"
)

type SendMessageRequest struct {
	MessagingProduct string      `json:"messaging_product"`
	To               string      `json:"to"`
	Type             MessageType `json:"type"`
}

type sendMessageErrorResponse struct {
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type WhatsappInterface interface {
	SendTextMessage(to string, textMessage SendTextMessage) (err error)
	SendMediaMessage(to string, mediaType MediaType, mediaMessage SendMediaMessage) (err error)
	SendInteractiveMessage(to string, interactiveMessage SendInteractiveMessage) (err error)
	SendTemplateMessage(to string, interactiveMessage SendTemplateMessage) (err error)
	ReadMedia(mediaID string) (body []byte, mimeType string, err error)
	ReceiveMessage(receiveMessage *ReceiveMessage, parsed func(*MessageParsed)) (err error)
	SetAsRead(messageID string) (err error)
}

type whatsapp struct {
	ID string
}

func Connect(id string) WhatsappInterface {
	return &whatsapp{
		ID: id,
	}
}

func (whatsapp *whatsapp) apiRequest(method string, path string, requestBody *[]byte) (responseBody []byte, err error) {
	body := &bytes.Reader{}
	if requestBody != nil {
		body = bytes.NewReader(*requestBody)
	}
	client := http.Client{}
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/%s", whatsapp.ID, path)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return responseBody, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WHATSAPP_KEY")))
	response, err := client.Do(request)
	if err != nil {
		return responseBody, err
	}
	defer response.Body.Close()
	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return responseBody, err
	}
	var sendMessageErrorResponse *sendMessageErrorResponse
	err = json.Unmarshal(responseBody, &sendMessageErrorResponse)
	if err != nil {
		return responseBody, err
	}
	if sendMessageErrorResponse.Error != nil {
		return responseBody, fmt.Errorf("message from facebook: error code %v - %s", sendMessageErrorResponse.Error.Code, sendMessageErrorResponse.Error.Message)
	}
	return responseBody, err
}

func (whatsapp *whatsapp) apiMediaDownloadRequest(path string, mimeType string) (responseBody []byte, err error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return responseBody, err
	}
	request.Header.Add("Content-Type", mimeType)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WHATSAPP_KEY")))
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
