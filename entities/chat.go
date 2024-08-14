package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type ChatMessage struct {
	ID          uuid.UUID       `json:"id" binding:"required"`
	SessionID   uuid.UUID       `json:"session" binding:"required"`
	Name        string          `json:"name" binding:"required"`
	Avatar      *string         `json:"avatar,omitempty"`
	MessageType string          `json:"message_type" binding:"required"`
	Body        *string         `json:"body,omitempty"`
	MediaID     *string         `json:"media_id,omitempty"`
	ActionID    *string         `json:"action_id,omitempty"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (ChatMessage) TableName() string {
	return "manager.chat__messages"
}

func (chatMessage *ChatMessage) BeforeCreate(tx *gorm.DB) (err error) {
	chatMessage.ID = uuid.New()
	chatMessage.CreatedAt = time.Now()
	return nil
}

type ChatMessageRepository interface {
	Insert(chatMessageData *ChatMessage) (insertedChatMessageData *ChatMessage, err error)
	Read(id uuid.UUID) (chatMessage *ChatMessage, err error)
	ReadAll(page int, pageSize int) (chatMessage *[]ChatMessage, pagination *utils.Pagination, err error)
	Update(chatMessageData *ChatMessage) (updatedChatMessage *ChatMessage, err error)
	Delete(id uuid.UUID) (err error)
}

type ChatSession struct {
	ID          uuid.UUID       `json:"id" binding:"required"`
	Channel     string          `json:"channel" binding:"required"`
	PhoneNumber *int            `json:"phone_number,omitempty"`
	Email       *string         `json:"email,omitempty"`
	Messages    []ChatMessage   `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"messages"`
	Status      bool            `gorm:"default:true" json:"status"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (ChatSession) TableName() string {
	return "manager.chat__session"
}

func (chatSession *ChatSession) BeforeCreate(tx *gorm.DB) (err error) {
	chatSession.ID = uuid.New()
	chatSession.CreatedAt = time.Now()
	chatSession.UpdatedAt = time.Now()
	return nil
}

func (chatSession *ChatSession) BeforeUpdate(tx *gorm.DB) (err error) {
	chatSession.UpdatedAt = time.Now()
	return nil
}

type ChatSessionRepository interface {
	Insert(chatSessionData *ChatSession) (insertedChatSessionData *ChatSession, err error)
	Read(id uuid.UUID) (chatSession *ChatSession, err error)
	ReadByPhoneNumber(phoneNumber int) (chatSession *ChatSession, err error)
	ReadAll(page int, pageSize int) (chatSession *[]ChatSession, pagination *utils.Pagination, err error)
	Update(chatSessionData *ChatSession) (updatedChatSession *ChatSession, err error)
	Delete(id uuid.UUID) (err error)
}
