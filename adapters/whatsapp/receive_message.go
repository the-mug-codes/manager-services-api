package whatsapp

type contact struct {
	WaID    string `json:"wa_id" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
	Profile struct {
		Name string `json:"name" binding:"required"`
	} `json:"profile" binding:"required"`
}

type whatsappError struct {
	Code      string `json:"code" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Message   string `json:"message" binding:"required"`
	ErrorData struct {
		Details string `json:"details" binding:"required"`
	} `json:"error_data" binding:"required"`
}

type media struct {
	ID       string  `json:"id" binding:"required"`
	MimeType string  `json:"mime_type" binding:"required"`
	Caption  *string `json:"caption,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Sha256   *string `json:"sha256,omitempty"`
}

type interactive struct {
	Type struct {
		ButtonReply *struct {
			ID    string `json:"id" binding:"required"`
			Title string `json:"title" binding:"required"`
		} `json:"button_reply,omitempty"`
		ListReply *struct {
			ID          string  `json:"id" binding:"required"`
			Title       string  `json:"title" binding:"required"`
			Description *string `json:"description,omitempty"`
		} `json:"list_reply,omitempty"`
	} `json:"type" binding:"required"`
}

type messages struct {
	ID       string      `json:"id" binding:"required"`
	From     string      `json:"from" binding:"required"`
	Type     MessageType `json:"type" binding:"required"`
	Audio    *media      `json:"audio,omitempty"`
	Document *media      `json:"document,omitempty"`
	Image    *media      `json:"image,omitempty"`
	Sticker  *media      `json:"sticker,omitempty"`
	Video    *media      `json:"video,omitempty"`
	Text     *struct {
		Body string `json:"body" binding:"required"`
	} `json:"text,omitempty"`
	Interactive *interactive   `json:"interactive,omitempty"`
	Errors      *whatsappError `json:"errors,omitempty"`
	Timestamp   string         `json:"timestamp" binding:"required"`
}

type statuses struct {
	ID          string `json:"id" binding:"required"`
	RecipientID string `json:"recipient_id" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Timestamp   int    `json:"timestamp" binding:"required"`
}

type value struct {
	MessagingProduct string `json:"messaging_product" binding:"required"`
	Metadata         struct {
		DisplayPhoneNumber string `json:"display_phone_number" binding:"required"`
		PhoneNumberID      string `json:"phone_number_id" binding:"required"`
	} `json:"metadata" binding:"required"`
	Contacts *[]contact  `json:"contacts,omitempty"`
	Errors   *[]error    `json:"errors,omitempty"`
	Messages *[]messages `json:"messages,omitempty"`
	Statuses *[]statuses `json:"statuses,omitempty"`
}

type changes struct {
	Field string `json:"field" binding:"required"`
	Value value  `json:"value" binding:"required"`
}

type entry struct {
	ID      string    `json:"id" binding:"required"`
	Changes []changes `json:"changes" binding:"required"`
}

type ReceiveMessage struct {
	Object string  `json:"object" binding:"required"`
	Entry  []entry `json:"entry" binding:"required"`
}

type MessageParsed struct {
	ID                 string  `json:"id" binding:"required"`
	Account            string  `json:"account" binding:"required"`
	DisplayPhoneNumber string  `json:"display_phone_number" binding:"required"`
	PhoneNumberID      *string `json:"phone_number_id" binding:"required"`
	ContactID          *string `json:"contact_id,omitempty"`
	ContactDisplayName *string `json:"contact_display_name,omitempty"`
	ContactPhoneNumber *string `json:"contact_phone_number,omitempty"`
	Status             *string `json:"status,omitempty"`
	MessageType        *string `json:"message_type,omitempty"`
	Message            *string `json:"message,omitempty"`
	MediaID            *string `json:"media_id,omitempty"`
	Filename           *string `json:"filename,omitempty"`
	MimeType           *string `json:"mime_type,omitempty"`
	ActionID           *string `json:"action_id,omitempty"`
}

func (whatsapp *whatsapp) ReceiveMessage(receiveMessage *ReceiveMessage, parsed func(*MessageParsed)) (err error) {
	var id string
	var account string
	var displayPhoneNumber string
	var phoneNumberID *string
	var contactID *string
	var contactDisplayName *string
	var contactPhoneNumber *string
	var status *string
	var message *string
	var messageType *string
	var mediaID *string
	var filename *string
	var mimeType *string
	var actionID *string
	if len(receiveMessage.Entry) == 0 {
		return err
	}
	for _, entry := range receiveMessage.Entry {
		account = entry.ID
		for _, changes := range entry.Changes {
			displayPhoneNumber = changes.Value.Metadata.DisplayPhoneNumber
			phoneNumberID = &changes.Value.Metadata.PhoneNumberID
			if changes.Value.Contacts != nil {
				contactsList := *changes.Value.Contacts
				contactID = &contactsList[0].UserID
				contactDisplayName = &contactsList[0].Profile.Name
				contactPhoneNumber = &contactsList[0].WaID
			}
			if changes.Value.Statuses != nil {
				statusList := *changes.Value.Statuses
				status = &statusList[0].Status
			}
			if changes.Value.Messages != nil {
				messagesList := *changes.Value.Messages
				messageType = (*string)(&messagesList[0].Type)
				id = messagesList[0].ID
				switch messagesList[0].Type {
				case MessageTypeText:
					message = &messagesList[0].Text.Body
				case MessageTypeAudio:
					mediaID = &messagesList[0].Audio.ID
					mimeType = &messagesList[0].Audio.MimeType
					message = messagesList[0].Audio.Caption
					filename = messagesList[0].Audio.Filename
				case MessageTypeDocument:
					mediaID = &messagesList[0].Document.ID
					mimeType = &messagesList[0].Document.MimeType
					message = messagesList[0].Document.Caption
					filename = messagesList[0].Document.Filename
				case MessageTypeImage:
					mediaID = &messagesList[0].Image.ID
					mimeType = &messagesList[0].Image.MimeType
					message = messagesList[0].Image.Caption
					filename = messagesList[0].Image.Filename
				case MessageTypeSticker:
					mediaID = &messagesList[0].Sticker.ID
					mimeType = &messagesList[0].Sticker.MimeType
					message = messagesList[0].Sticker.Caption
					filename = messagesList[0].Sticker.Filename
				case MessageTypeVideo:
					mediaID = &messagesList[0].Video.ID
					mimeType = &messagesList[0].Video.MimeType
					message = messagesList[0].Video.Caption
					filename = messagesList[0].Video.Filename
				case MessageTypeInteractive:
					if messagesList[0].Interactive.Type.ButtonReply != nil {
						message = &messagesList[0].Interactive.Type.ButtonReply.Title
						actionID = &messagesList[0].Interactive.Type.ButtonReply.ID
					}
					if messagesList[0].Interactive.Type.ListReply != nil {
						message = &messagesList[0].Interactive.Type.ListReply.Title
						actionID = &messagesList[0].Interactive.Type.ListReply.ID
					}
				}
			}
			go parsed(&MessageParsed{
				ID:                 id,
				Account:            account,
				DisplayPhoneNumber: displayPhoneNumber,
				PhoneNumberID:      phoneNumberID,
				ContactID:          contactID,
				ContactDisplayName: contactDisplayName,
				ContactPhoneNumber: contactPhoneNumber,
				Status:             status,
				MessageType:        messageType,
				Message:            message,
				MediaID:            mediaID,
				Filename:           filename,
				MimeType:           mimeType,
				ActionID:           actionID,
			})
		}
	}
	return err
}
