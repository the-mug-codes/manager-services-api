package messagebird

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Contact struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Attributes  struct {
		ID uuid.UUID `json:"id"`
	} `json:"attributes"`
	Status string `json:"status"`
}

type ContactList struct {
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	Count      int       `json:"count"`
	TotalCount int       `json:"totalCount"`
	Items      []Contact `json:"items"`
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetAllContacts() (contacts *ContactList, err error) {
	responseBody, err := messageBird.apiRequest("GET", "contacts", "/v2/contacts/", nil, nil)
	if err != nil {
		return contacts, err
	}
	err = json.Unmarshal(responseBody, &contacts)
	if err != nil {
		return contacts, err
	}
	return contacts, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetContact(id string) (contact *Contact, err error) {
	path := fmt.Sprintf("/v2/contacts/%s", id)
	responseBody, err := messageBird.apiRequest("GET", "contacts", path, nil, nil)
	if err != nil {
		return contact, err
	}
	err = json.Unmarshal(responseBody, &contact)
	if err != nil {
		return contact, err
	}
	return contact, err
}
