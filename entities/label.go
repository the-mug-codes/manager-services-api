package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type Label struct {
	ID        uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	Name      string          `gorm:"index;unique" json:"name" binding:"required"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func (Label) TableName() string {
	return "manager.labels"
}

func (label *Label) BeforeCreate(tx *gorm.DB) (err error) {
	if label.ID == uuid.Nil {
		label.ID = uuid.New()
	}
	label.CreatedAt = time.Now()
	label.UpdatedAt = time.Now()
	return
}

func (label *Label) BeforeUpdate(tx *gorm.DB) (err error) {
	label.UpdatedAt = time.Now()
	return
}

type LabelRepository interface {
	Insert(labelData *Label) (insertedLabel *Label, err error)
	Read(id uuid.UUID) (label *Label, err error)
	ReadAll(page int, pageSize int) (labels *[]Label, pagination *utils.Pagination, err error)
	Update(labelData *Label) (updatedLabel *Label, err error)
	Delete(id uuid.UUID) (err error)
}
