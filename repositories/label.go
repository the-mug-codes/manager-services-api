package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	"gorm.io/gorm"
)

type labelRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func Label(context *gin.Context) entity.LabelRepository {
	return &labelRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *labelRepository) Insert(labelData *entity.Label) (insertedLabel *entity.Label, err error) {
	err = connection.database.Create(&labelData).Error
	if err != nil {
		return insertedLabel, err
	}
	err = connection.database.First(&insertedLabel, labelData.ID).Error
	if err != nil {
		return insertedLabel, err
	}
	return insertedLabel, err
}

func (connection *labelRepository) ReadAll(page int, pageSize int) (labels *[]entity.Label, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Find(&entity.Label{}).Count(&totalRegisters).Error
	if err != nil {
		return labels, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Find(&labels).Error
	if err != nil {
		return labels, pagination, err
	}
	return labels, pagination, err
}

func (connection *labelRepository) Read(id uuid.UUID) (label *entity.Label, err error) {
	err = connection.database.First(&label, id).Error
	if err != nil {
		return label, err
	}
	return label, err
}

func (connection *labelRepository) Update(labelData *entity.Label) (updatedLabel *entity.Label, err error) {
	err = connection.database.Omit("CreatedAt").Save(&labelData).Error
	if err != nil {
		return updatedLabel, err
	}
	err = connection.database.First(&updatedLabel, labelData.ID).Error
	if err != nil {
		return updatedLabel, err
	}
	return updatedLabel, err
}

func (connection *labelRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.Label{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
