package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type HelpDeskTicketMessage struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TicketID  uuid.UUID `json:"ticket_id" binding:"required"`
	Text      string    `json:"description" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type HelpDeskTicket struct {
	ID        uuid.UUID               `gorm:"type:uuid;primary_key" json:"id"`
	Title     string                  `json:"title" binding:"required"`
	Level     string                  `gorm:"default:medium" json:"level" binding:"required"`
	Code      int                     `json:"code" binding:"required"`
	Labels    []*Label                `gorm:"many2many:manager.helpdesk_tickets__labels;constraint:OnDelete:CASCADE" json:"labels,omitempty"`
	PersonID  uuid.UUID               `gorm:"type:uuid" json:"person_id"`
	Person    Person                  `json:"person,omitempty"`
	ProjectID *uuid.UUID              `gorm:"type:uuid" json:"project_id,omitempty"`
	Project   *Project                `json:"project,omitempty"`
	Messages  []HelpDeskTicketMessage `gorm:"many2many:manager.helpdesk_tickets__messages;constraint:OnDelete:CASCADE" json:"messages"`
	DueDate   time.Time               `json:"due_date" binding:"required"`
	Source    string                  `json:"source" binding:"required"`
	SourceID  string                  `json:"source_id" binding:"required"`
	Status    bool                    `gorm:"default:true" json:"status"`
	CreatedAt time.Time               `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time               `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt         `gorm:"index" json:"deleted_at,omitempty"`
}

func (HelpDeskTicketMessage) TableName() string {
	return "manager.helpdesk_tickets__messages"
}

func (HelpDeskTicket) TableName() string {
	return "manager.helpdesk_tickets"
}

func (ticketMessage *HelpDeskTicketMessage) BeforeCreate(tx *gorm.DB) (err error) {
	ticketMessage.ID = uuid.New()
	ticketMessage.CreatedAt = time.Now()
	return nil
}

func (ticket *HelpDeskTicket) BeforeCreate(tx *gorm.DB) (err error) {
	ticket.ID = uuid.New()
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	return nil
}

func (ticket *HelpDeskTicket) BeforeUpdate(tx *gorm.DB) (err error) {
	ticket.UpdatedAt = time.Now()
	return nil
}

type HelpDeskTicketMessageRepository interface {
	Insert(ticketMessageData *HelpDeskTicketMessage) (insertedTicketMessage *HelpDeskTicketMessage, err error)
	Read(id uuid.UUID) (ticketMessage *HelpDeskTicketMessage, err error)
	ReadAll(page int, pageSize int) (ticketMessages *[]HelpDeskTicketMessage, pagination *utils.Pagination, err error)
	ReadAllByTicket(ticketID uuid.UUID, page int, pageSize int) (ticketMessages *[]HelpDeskTicketMessage, pagination *utils.Pagination, err error)
}

type HelpDeskTicketRepository interface {
	Insert(ticketData *HelpDeskTicket) (insertedTicket *HelpDeskTicket, err error)
	Read(id uuid.UUID) (ticket *HelpDeskTicket, err error)
	ReadAll(page int, pageSize int) (tickets *[]HelpDeskTicket, pagination *utils.Pagination, err error)
	Update(ticketData *HelpDeskTicket) (updatedTicket *HelpDeskTicket, err error)
	Delete(id uuid.UUID) (err error)
}
