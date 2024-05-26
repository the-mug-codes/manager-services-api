package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type productCategoryRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func ProductCategory(context *gin.Context) entity.ProductCategoryRepository {
	return &productCategoryRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *productCategoryRepository) Insert(productCategoryData *entity.ProductCategory) (insertedProductCategory *entity.ProductCategory, err error) {
	err = connection.database.Create(&productCategoryData).Error
	if err != nil {
		return insertedProductCategory, err
	}
	err = connection.database.Preload("Products").First(&insertedProductCategory, productCategoryData.ID).Error
	if err != nil {
		return insertedProductCategory, err
	}
	return insertedProductCategory, err
}

func (connection *productCategoryRepository) ReadAll(page int, pageSize int) (productCategorys *[]entity.ProductCategory, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Find(&entity.ProductCategory{}).Count(&totalRegisters).Error
	if err != nil {
		return productCategorys, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Preload("Products").Find(&productCategorys).Error
	if err != nil {
		return productCategorys, pagination, err
	}
	return productCategorys, pagination, err
}

func (connection *productCategoryRepository) Read(id uuid.UUID) (productCategory *entity.ProductCategory, err error) {
	err = connection.database.Preload("Products").First(&productCategory, id).Error
	if err != nil {
		return productCategory, err
	}
	return productCategory, err
}

func (connection *productCategoryRepository) Update(productCategoryData *entity.ProductCategory) (updatedProductCategory *entity.ProductCategory, err error) {
	err = connection.database.Preload("Products").First(&updatedProductCategory, productCategoryData.ID).Error
	if err != nil {
		return updatedProductCategory, err
	}
	return updatedProductCategory, err
}

func (connection *productCategoryRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.ProductCategory{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
