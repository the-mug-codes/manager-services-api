package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type Project struct {
	ID           uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	Kind         string          `json:"kind" binding:"required"`
	Title        string          `gorm:"index;unique" json:"title" binding:"required"`
	Labels       []*Label        `gorm:"many2many:manager.projects__labels;constraint:OnDelete:CASCADE" json:"labels"`
	PersonID     uuid.UUID       `gorm:"type:uuid" json:"person_id"`
	Person       *Person         `json:"person,omitempty"`
	Requirements pq.StringArray  `gorm:"type:text[]" json:"requirements" binding:"required"`
	Content      string          `json:"content" binding:"required"`
	Comments     *string         `json:"comments"`
	Deliveries   pq.StringArray  `gorm:"type:text[]" json:"deliveries"`
	DueDate      *time.Time      `json:"due_date"`
	CreatedAt    time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Project) TableName() string {
	return "manager.projects"
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	project.ID = uuid.New()
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	return nil
}

func (project *Project) BeforeUpdate(tx *gorm.DB) (err error) {
	project.UpdatedAt = time.Now()
	return nil
}

type ProjectRepository interface {
	Insert(projectData *Project) (insertedProject *Project, err error)
	Read(id uuid.UUID) (project *Project, err error)
	ReadAll(page int, pageSize int) (projects *[]Project, pagination *utils.Pagination, err error)
	Update(projectData *Project) (updatedProject *Project, err error)
	Delete(id uuid.UUID) (err error)
}
