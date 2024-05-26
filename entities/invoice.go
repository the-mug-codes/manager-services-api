package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InvoiceItem struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	InvoiceID   uuid.UUID       `gorm:"type:uuid" json:"invoice_id"`
	ProductID   uuid.UUID       `gorm:"type:uuid" json:"product_id"`
	Product     Product         `json:"product"`
	PriceValue  float64         `gorm:"type:decimal(10,2);default:0" json:"price_value" binding:"required"`
	Quantity    float64         `gorm:"type:decimal(10,2);default:1" json:"quantity" binding:"required"`
	TotalAmount float64         `gorm:"type:decimal(10,2);default:0" json:"total_amount" binding:"required"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Invoice struct {
	ID                  uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	Code                string          `json:"code" binding:"required"`
	ContractID          *uuid.UUID      `json:"contract_id,omitempty"`
	PersonID            uuid.UUID       `gorm:"type:uuid" json:"person_id"`
	Person              Person          `json:"person"`
	Items               []*InvoiceItem  `gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE" json:"items"`
	DueDate             time.Time       `json:"due_date" binding:"required"`
	AmountValue         float64         `gorm:"default:0" json:"amount_value" binding:"required"`
	DiscountValue       float64         `gorm:"default:0" json:"discount_value" binding:"required"`
	ExtraAmountValue    float64         `gorm:"default:0" json:"extra_amount_value" binding:"required"`
	TotalAmountValue    float64         `gorm:"default:0" json:"total_amount_value" binding:"required"`
	InstallmentNumber   *int            `json:"installment_number,omitempty"`
	PixCopyPaste        *string         `json:"pix_copy_paste,omitempty"`
	PixQrCode           *string         `json:"pix_qr_code,omitempty"`
	BankSlipCode        *string         `json:"bank_slip_code,omitempty"`
	BankSlipDocument    *string         `json:"bank_slip_document,omitempty"`
	BankSlipInputCode   *string         `json:"bank_slip_input_code,omitempty"`
	PaidAt              *time.Time      `json:"paid_at,omitempty"`
	PaymentMethod       *string         `json:"payment_method,omitempty"`
	PaymentComments     *string         `json:"payment_comments,omitempty"`
	Comments            *string         `json:"comments,omitempty"`
	DiscountComments    *string         `json:"discount_comments,omitempty"`
	ExtraAmountComments *string         `json:"extra_amount_comments,omitempty"`
	WarningComments     *string         `json:"warning_comments,omitempty"`
	Status              bool            `gorm:"default:true" json:"status"`
	LateInvoices        *bool           `gorm:"default:false" json:"late_invoices"`
	CreatedAt           time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt           *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Invoice) ParseDate(date time.Time) string {
	return date.Format("02/01/2006")
}

func (Invoice) ParseMoney(amount float64) string {
	return fmt.Sprintf("R$ %.2f", amount)
}

func (invoice Invoice) GetQrCode() (base64 string) {
	if invoice.PixQrCode != nil {
		base64 = *invoice.PixQrCode
		return base64[22:]
	}
	return base64
}

func (Invoice) TableName() string {
	return "manager.finance__invoices"
}

func (InvoiceItem) TableName() string {
	return "manager.finance__invoices_items"
}

func (InvoiceItem) ParseMoney(amount float64) string {
	return fmt.Sprintf("R$ %.2f", amount)
}

func (invoice *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	invoice.ID = uuid.New()
	invoice.CreatedAt = time.Now()
	return nil
}

func (invoiceItem *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	invoiceItem.ID = uuid.New()
	invoiceItem.CreatedAt = time.Now()
	return nil
}

func (invoice *Invoice) BeforeUpdate(tx *gorm.DB) (err error) {
	invoice.UpdatedAt = time.Now()
	return nil
}

func (invoice *Invoice) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("invoice_id = ?", invoice.ID).Delete(&InvoiceItem{})
	return nil
}

type InvoiceRepository interface {
	Insert(invoice *Invoice) (insertedInvoice *Invoice, err error)
	Read(id uuid.UUID) (invoice *Invoice, err error)
	ReadAll(page int, pageSize int) (invoices *[]Invoice, pagination *utils.Pagination, err error)
	Update(invoice *Invoice) (updatedInvoices *Invoice, err error)
	Delete(id uuid.UUID) (err error)
}

type InvoiceHtml interface {
	Generate(templatePath string, data Invoice) (html *string, err error)
}

type InvoicePDF interface {
	GenerateFile(documentId string, documentContent string, documentName string, isDocument bool, noMargin bool, pageSize string, orientation string) (documentRetrievePath string, err error)
	GenerateBase64(documentPath string, delete bool, formatted bool) (base64 *string, err error)
	GenerateBinary(documentPath string, delete bool) (binary *[]byte, err error)
}

type InvoiceWhatsApp interface {
	SendMessage(to string, from string, messageType string, contentMessage MessageContent, replyTo *string) (messageSent *NewMessageCreated, err error)
}

type InvoiceEmail interface {
	SendEmailMessage(to string, name string, subject string, text string, html *string, attachment *[]EmailAttachment) (err error)
}
