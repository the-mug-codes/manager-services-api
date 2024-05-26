package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type Contract struct {
	ID                uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	ContractModel     string          `json:"contract_model" binding:"required"`
	Labels            []*Label        `gorm:"many2many:manager.contracts__labels;constraint:OnDelete:CASCADE" json:"labels,omitempty"`
	PersonID          uuid.UUID       `gorm:"type:uuid" json:"person_id"`
	Person            Person          `json:"person"`
	ProjectID         uuid.UUID       `gorm:"type:uuid" json:"project_id"`
	Project           Project         `json:"project"`
	Description       string          `json:"description" binding:"required"`
	Comments          string          `json:"comments" binding:"required"`
	AmountValue       float64         `gorm:"default:0" json:"amount_value" binding:"required"`
	DiscountValue     float64         `gorm:"default:0" json:"discount_value" binding:"required"`
	ExtraAmountValue  float64         `gorm:"default:0" json:"extra_amount_value" binding:"required"`
	TotalAmountValue  float64         `gorm:"default:0" json:"total_amount_value" binding:"required"`
	InstallmentNumber *int            `json:"installment_number,omitempty"`
	DueDate           *time.Time      `json:"due_date" binding:"required"`
	SignedAt          *time.Time      `json:"signed_at,omitempty"`
	Document          string          `json:"document,omitempty"`
	ResellerID        *uuid.UUID      `json:"reseller_id"`
	Products          []*Product      `gorm:"many2many:manager.contracts__products;constraint:OnDelete:CASCADE" json:"products,omitempty"`
	Invoices          []*Invoice      `gorm:"many2many:manager.contracts__invoices;constraint:OnDelete:CASCADE" json:"invoices,omitempty"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Contract) TableName() string {
	return "manager.contracts"
}

func (contract *Contract) BeforeCreate(tx *gorm.DB) (err error) {
	contract.ID = uuid.New()
	contract.CreatedAt = time.Now()
	contract.UpdatedAt = time.Now()
	return nil
}

func (contract *Contract) BeforeUpdate(tx *gorm.DB) (err error) {
	contract.UpdatedAt = time.Now()
	return nil
}

type ContractRepository interface {
	Insert(contractData *Contract) (insertedContract *Contract, err error)
	Read(id uuid.UUID) (contract *Contract, err error)
	ReadAll(page int, pageSize int) (contracts *[]Contract, pagination *utils.Pagination, err error)
	Update(contractData *Contract) (updatedContract *Contract, err error)
	Delete(id uuid.UUID) (err error)
}
