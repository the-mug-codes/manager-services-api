package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	"gorm.io/gorm"
)

type productRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func Product(context *gin.Context) entity.ProductRepository {
	return &productRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *productRepository) Insert(productData *entity.Product) (insertedProduct *entity.Product, err error) {
	err = connection.database.Create(&productData).Error
	if err != nil {
		return insertedProduct, err
	}
	err = connection.database.Preload("Prices").Preload("Categories").First(&insertedProduct, productData.ID).Error
	if err != nil {
		return insertedProduct, err
	}
	return insertedProduct, err
}

func (connection *productRepository) ReadAll(page int, pageSize int) (products *[]entity.Product, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Find(&entity.Product{}).Count(&totalRegisters).Error
	if err != nil {
		return products, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Preload("Prices").Preload("Categories").Find(&products).Error
	if err != nil {
		return products, pagination, err
	}
	return products, pagination, err
}

func (connection *productRepository) Read(id uuid.UUID) (product *entity.Product, err error) {
	err = connection.database.Preload("Prices").Preload("Categories").First(&product, id).Error
	if err != nil {
		return product, err
	}
	return product, err
}

func (connection *productRepository) Update(productData *entity.Product) (updatedProduct *entity.Product, err error) {
	err = connection.database.Preload("Prices").Preload("Categories").First(&updatedProduct, productData.ID).Error
	if err != nil {
		return updatedProduct, err
	}
	return updatedProduct, err
}

func (connection *productRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
