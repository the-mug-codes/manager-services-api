package messagebird

import (
	"encoding/json"
	"fmt"
	"time"
)

type Conversation struct {
	ID                string    `json:"id"`
	ContactID         string    `json:"contactId"`
	LastUsedPlatform  string    `json:"lastUsedPlatformId"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"createdDatetime"`
	UpdatedAt         time.Time `json:"updatedDatetime"`
	LastInteractionAt time.Time `json:"lastReceivedDatetime"`
	Messages          struct {
		TotalCount int `json:"totalCount"`
	} `json:"messages"`
}

type ConversationList struct {
	Offset     int            `json:"offset"`
	Limit      int            `json:"limit"`
	Count      int            `json:"count"`
	TotalCount int            `json:"totalCount"`
	Items      []Conversation `json:"items"`
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

type ConversationMessages struct {
	ID        string         `json:"id"`
	Platform  string         `json:"platform"`
	To        string         `json:"to"`
	From      string         `json:"from"`
	Direction string         `json:"direction"`
	Status    string         `json:"status"`
	Type      string         `json:"type"`
	CreatedAt time.Time      `json:"createdDatetime"`
	UpdatedAt time.Time      `json:"updatedDatetime"`
	Origin    string         `json:"origin"`
	Content   MessageContent `json:"content"`
}

type ConversationMessagesList struct {
	Offset     int                    `json:"offset"`
	Limit      int                    `json:"limit"`
	Count      int                    `json:"count"`
	TotalCount int                    `json:"totalCount"`
	Items      []ConversationMessages `json:"items"`
}

type NewMessage[MessageContent any] struct {
	To      string         `json:"to"`
	From    string         `json:"from"`
	Type    string         `json:"type"`
	Content MessageContent `json:"content"`
	ReplyTo *struct {
		ID string `json:"id"`
	} `json:"replyTo"`
}

type NewMessageCreated struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetAllConversations() (conversations *ConversationList, err error) {
	queryParams := &[]map[string]string{
		{
			"status": "archived;active",
		},
	}
	responseBody, err := messageBird.apiRequest("GET", "conversations", "/v1/conversations", nil, queryParams)
	if err != nil {
		return conversations, err
	}
	err = json.Unmarshal(responseBody, &conversations)
	if err != nil {
		return conversations, err
	}
	return conversations, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetConversation(id string) (conversation *Conversation, err error) {
	path := fmt.Sprintf("/v1/conversations/%s", id)
	responseBody, err := messageBird.apiRequest("GET", "conversations", path, nil, nil)
	if err != nil {
		return conversation, err
	}
	err = json.Unmarshal(responseBody, &conversation)
	if err != nil {
		return conversation, err
	}
	return conversation, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetConversationMessages(id string) (conversationMessages *ConversationMessagesList, err error) {
	path := fmt.Sprintf("/v1/conversations/%s/messages", id)
	responseBody, err := messageBird.apiRequest("GET", "conversations", path, nil, nil)
	if err != nil {
		return conversationMessages, err
	}
	err = json.Unmarshal(responseBody, &conversationMessages)
	if err != nil {
		return conversationMessages, err
	}
	return conversationMessages, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) SendMessage(to string, from string, messageType string, messageContent MessageContent, replyTo *string) (messageSent *NewMessageCreated, err error) {
	newMessageBody := NewMessage[MessageContent]{
		To:      to,
		Type:    messageType,
		From:    from,
		Content: messageContent,
	}
	if replyTo != nil {
		newMessageBody.ReplyTo = &struct {
			ID string "json:\"id\""
		}{
			ID: *replyTo,
		}
	}
	requestBody, err := json.Marshal(newMessageBody)
	if err != nil {
		return messageSent, err
	}
	responseBody, err := messageBird.apiRequest("POST", "conversations", "/v1/send", &requestBody, nil)
	if err != nil {
		return messageSent, err
	}
	err = json.Unmarshal(responseBody, &messageSent)
	if err != nil {
		return messageSent, err
	}
	return messageSent, err
}
