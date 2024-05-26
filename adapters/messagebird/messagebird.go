package messagebird

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type MessageBirdInterface[MessageContent any, NewMessageCreated any] interface {
	GetAllContacts() (contacts *ContactList, err error)
	GetContact(id string) (contact *Contact, err error)
	GetAllConversations() (conversations *ConversationList, err error)
	GetConversation(id string) (conversation *Conversation, err error)
	GetConversationMessages(id string) (conversationMessages *ConversationMessagesList, err error)
	SendMessage(to string, from string, messageType string, contentMessage MessageContent, replyTo *string) (messageSent *NewMessageCreated, err error)
	GetAllCalls() (calls *CallList, err error)
	GetCall(id string) (call *CallItem, err error)
	GetCallRecording(id string) (recording *RecordingList, err error)
	GetCallRecordingFile(id string, format string) (recordingFile *[]byte, err error)
	GetCallRecordingTranscription(id string) (recordingTranscription *string, err error)
}

type messageBird[MessageContent any, NewMessageCreated any] struct{}

func Connect[MessageContent any, NewMessageCreated any]() MessageBirdInterface[MessageContent, NewMessageCreated] {
	return &messageBird[MessageContent, NewMessageCreated]{}
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) apiRequest(method string, service string, path string, requestBody *[]byte, queryParams *[]map[string]string) (responseBody []byte, err error) {
	body := &bytes.Reader{}
	if requestBody != nil {
		body = bytes.NewReader(*requestBody)
	}
	client := http.Client{}
	url := fmt.Sprintf("https://%s.messagebird.com%s", service, path)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return responseBody, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("AccessKey %s", os.Getenv("MESSAGE_BIRD_KEY")))
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
