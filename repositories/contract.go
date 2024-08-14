package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	"gorm.io/gorm"
)

type contractRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func Contract(context *gin.Context) entity.ContractRepository {
	return &contractRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *contractRepository) Insert(contractData *entity.Contract) (insertedContract *entity.Contract, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &contractData.PersonID) != nil {
		return insertedContract, err
	}
	err = connection.database.Create(&contractData).Error
	if err != nil {
		return insertedContract, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Products").Preload("Invoices").First(&insertedContract, contractData.ID).Error
	if err != nil {
		return insertedContract, err
	}
	return insertedContract, err
}

func (connection *contractRepository) ReadAll(page int, pageSize int) (contracts *[]entity.Contract, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Find(&entity.Contract{}).Count(&totalRegisters).Error
	if err != nil {
		return contracts, pagination, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Preload("Labels").Preload("Products").Preload("Invoices").Find(&contracts).Error
	if err != nil {
		return contracts, pagination, err
	}
	return contracts, pagination, err
}

func (connection *contractRepository) Read(id uuid.UUID) (contract *entity.Contract, err error) {
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Products").Preload("Invoices").First(&contract, id).Error
	if err != nil {
		return contract, err
	}
	return contract, err
}

func (connection *contractRepository) Update(contractData *entity.Contract) (updatedContract *entity.Contract, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &contractData.PersonID) != nil {
		return updatedContract, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Products").Preload("Invoices").First(&updatedContract, contractData.ID).Error
	if err != nil {
		return updatedContract, err
	}
	return updatedContract, err
}

func (connection *contractRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.Contract{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
