package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type ticketMessageRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func HelpDeskTicketMessage(context *gin.Context) entity.HelpDeskTicketMessageRepository {
	return &ticketMessageRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *ticketMessageRepository) Insert(ticketMessageData *entity.HelpDeskTicketMessage) (insertedTicketMessage *entity.HelpDeskTicketMessage, err error) {
	err = connection.database.Create(&ticketMessageData).Error
	if err != nil {
		return insertedTicketMessage, err
	}
	err = connection.database.First(&insertedTicketMessage, ticketMessageData.ID).Error
	if err != nil {
		return insertedTicketMessage, err
	}
	return insertedTicketMessage, err
}

func (connection *ticketMessageRepository) ReadAll(page int, pageSize int) (ticketMessages *[]entity.HelpDeskTicketMessage, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Find(&entity.HelpDeskTicketMessage{}).Count(&totalRegisters).Error
	if err != nil {
		return ticketMessages, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Find(&ticketMessages).Error
	if err != nil {
		return ticketMessages, pagination, err
	}
	return ticketMessages, pagination, err
}

func (connection *ticketMessageRepository) ReadAllByTicket(ticketID uuid.UUID, page int, pageSize int) (ticketMessages *[]entity.HelpDeskTicketMessage, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Find(&entity.HelpDeskTicketMessage{}).Count(&totalRegisters).Error
	if err != nil {
		return ticketMessages, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Find(&ticketMessages).Error
	if err != nil {
		return ticketMessages, pagination, err
	}
	return ticketMessages, pagination, err
}

func (connection *ticketMessageRepository) Read(id uuid.UUID) (ticketMessage *entity.HelpDeskTicketMessage, err error) {
	err = connection.database.First(&ticketMessage, id).Error
	if err != nil {
		return ticketMessage, err
	}
	return ticketMessage, err
}
