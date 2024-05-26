package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type invoiceRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func Invoice(context *gin.Context) entity.InvoiceRepository {
	return &invoiceRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *invoiceRepository) Insert(invoiceData *entity.Invoice) (insertedInvoice *entity.Invoice, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &invoiceData.PersonID) != nil {
		return insertedInvoice, err
	}
	err = connection.database.Create(&invoiceData).Error
	if err != nil {
		return insertedInvoice, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Items").First(&insertedInvoice, invoiceData.ID).Error
	if err != nil {
		return insertedInvoice, err
	}
	return insertedInvoice, err
}

func (connection *invoiceRepository) ReadAll(page int, pageSize int) (invoices *[]entity.Invoice, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Find(&entity.Invoice{}).Count(&totalRegisters).Error
	if err != nil {
		return invoices, pagination, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Preload("Items").Find(&invoices).Error
	if err != nil {
		return invoices, pagination, err
	}
	return invoices, pagination, err
}

func (connection *invoiceRepository) Read(id uuid.UUID) (invoice *entity.Invoice, err error) {
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Items").First(&invoice, id).Error
	if err != nil {
		return invoice, err
	}
	return invoice, err
}

func (connection *invoiceRepository) Update(invoiceData *entity.Invoice) (updatedInvoice *entity.Invoice, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &invoiceData.PersonID) != nil {
		return updatedInvoice, err
	}
	err = connection.database.Omit("Items").Where("id = ?", invoiceData.ID).Updates(&invoiceData).Error
	if err != nil {
		return updatedInvoice, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Items").First(&updatedInvoice, invoiceData.ID).Error
	if err != nil {
		return updatedInvoice, err
	}
	return updatedInvoice, err
}

func (connection *invoiceRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.Invoice{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
