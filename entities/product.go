package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type ProductPrice struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	Code        string          `json:"code" binding:"required"`
	Title       string          `json:"title"`
	Description string          `json:"description" binding:"required"`
	Value       float32         `gorm:"type:decimal(10,2);default:0" json:"value" binding:"required"`
	ProductID   uuid.UUID       `gorm:"type:uuid" json:"product_id"`
	Status      bool            `gorm:"default:true" json:"status"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type ProductCategory struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Products    []Product       `gorm:"many2many:manager.products__in_categories" json:"products,omitempty"`
	Status      bool            `gorm:"default:true" json:"status"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Product struct {
	ID          uuid.UUID         `gorm:"type:uuid;primary_key" json:"id"`
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description" binding:"required"`
	Prices      []ProductPrice    `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"prices"`
	Categories  []ProductCategory `gorm:"many2many:manager.products__in_categories" json:"categories,omitempty"`
	Status      bool              `gorm:"default:true" json:"status"`
	CreatedAt   time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt   `gorm:"index" json:"deleted_at,omitempty"`
}

func (ProductPrice) TableName() string {
	return "manager.products__prices"
}

func (ProductCategory) TableName() string {
	return "manager.products__categories"
}

func (Product) TableName() string {
	return "manager.products"
}

func (productPrice *ProductPrice) BeforeCreate(tx *gorm.DB) (err error) {
	productPrice.ID = uuid.New()
	productPrice.CreatedAt = time.Now()
	productPrice.UpdatedAt = time.Now()
	return
}

func (productPrice *ProductPrice) BeforeUpdate(tx *gorm.DB) (err error) {
	productPrice.UpdatedAt = time.Now()
	return nil
}

func (productCategory *ProductCategory) BeforeCreate(tx *gorm.DB) (err error) {
	productCategory.ID = uuid.New()
	productCategory.CreatedAt = time.Now()
	productCategory.UpdatedAt = time.Now()
	return
}

func (productCategory *ProductCategory) BeforeUpdate(tx *gorm.DB) (err error) {
	productCategory.UpdatedAt = time.Now()
	return nil
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	return
}

func (productCategory *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	productCategory.UpdatedAt = time.Now()
	return nil
}

type ProductCategoryRepository interface {
	Insert(productCategoryData *ProductCategory) (insertedProductCategory *ProductCategory, err error)
	Read(id uuid.UUID) (productCategory *ProductCategory, err error)
	ReadAll(page int, pageSize int) (productCategories *[]ProductCategory, pagination *utils.Pagination, err error)
	Update(productCategoryData *ProductCategory) (updatedProductCategory *ProductCategory, err error)
	Delete(id uuid.UUID) (err error)
}

type ProductRepository interface {
	Insert(productData *Product) (insertedProduct *Product, err error)
	Read(id uuid.UUID) (product *Product, err error)
	ReadAll(page int, pageSize int) (products *[]Product, pagination *utils.Pagination, err error)
	Update(productData *Product) (updatedProduct *Product, err error)
	Delete(id uuid.UUID) (err error)
}
