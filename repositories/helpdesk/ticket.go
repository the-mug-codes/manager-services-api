package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type ticketRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func HelpDeskTicket(context *gin.Context) entity.HelpDeskTicketRepository {
	return &ticketRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *ticketRepository) Insert(ticketData *entity.HelpDeskTicket) (insertedTicket *entity.HelpDeskTicket, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &ticketData.PersonID) != nil {
		return insertedTicket, err
	}
	err = connection.database.Create(&ticketData).Error
	if err != nil {
		return insertedTicket, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Messages").First(&insertedTicket, ticketData.ID).Error
	if err != nil {
		return insertedTicket, err
	}
	return insertedTicket, err
}

func (connection *ticketRepository) ReadAll(page int, pageSize int) (tickets *[]entity.HelpDeskTicket, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Find(&entity.HelpDeskTicket{}).Count(&totalRegisters).Error
	if err != nil {
		return tickets, pagination, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Preload("Labels").Preload("Messages").Find(&tickets).Error
	if err != nil {
		return tickets, pagination, err
	}
	return tickets, pagination, err
}

func (connection *ticketRepository) Read(id uuid.UUID) (ticket *entity.HelpDeskTicket, err error) {
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Messages").First(&ticket, id).Error
	if err != nil {
		return ticket, err
	}
	return ticket, err
}

func (connection *ticketRepository) Update(ticketData *entity.HelpDeskTicket) (updatedTicket *entity.HelpDeskTicket, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &ticketData.PersonID) != nil {
		return updatedTicket, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Messages").First(&updatedTicket, ticketData.ID).Error
	if err != nil {
		return updatedTicket, err
	}
	return updatedTicket, err
}

func (connection *ticketRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.HelpDeskTicket{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
