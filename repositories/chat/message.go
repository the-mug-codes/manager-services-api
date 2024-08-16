package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	"gorm.io/gorm"
)

type chatMessageRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func ChatMessage(context *gin.Context) entity.ChatMessageRepository {
	return &chatMessageRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *chatMessageRepository) Insert(chatMessageData *entity.ChatMessage) (insertedChatMessage *entity.ChatMessage, err error) {
	err = connection.database.Create(&chatMessageData).Error
	if err != nil {
		return insertedChatMessage, err
	}
	err = connection.database.First(&insertedChatMessage, chatMessageData.ID).Error
	if err != nil {
		return insertedChatMessage, err
	}
	return insertedChatMessage, err
}

func (connection *chatMessageRepository) ReadAll(page int, pageSize int) (chatMessages *[]entity.ChatMessage, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Find(&entity.ChatMessage{}).Count(&totalRegisters).Error
	if err != nil {
		return chatMessages, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Find(&chatMessages).Error
	if err != nil {
		return chatMessages, pagination, err
	}
	return chatMessages, pagination, err
}

func (connection *chatMessageRepository) ReadAllBySession(sessionID uuid.UUID, page int, pageSize int) (chatMessages *[]entity.ChatMessage, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Where("session_id = ?", sessionID).Find(&entity.ChatMessage{}).Count(&totalRegisters).Error
	if err != nil {
		return chatMessages, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Where("session_id = ?", sessionID).Find(&chatMessages).Error
	if err != nil {
		return chatMessages, pagination, err
	}
	return chatMessages, pagination, err
}

func (connection *chatMessageRepository) Read(id uuid.UUID) (chatMessage *entity.ChatMessage, err error) {
	err = connection.database.First(&chatMessage, id).Error
	if err != nil {
		return chatMessage, err
	}
	return chatMessage, err
}

func (connection *chatMessageRepository) Update(chatMessageData *entity.ChatMessage) (updatedChatMessage *entity.ChatMessage, err error) {
	err = connection.database.First(&updatedChatMessage, chatMessageData.ID).Error
	if err != nil {
		return updatedChatMessage, err
	}
	return updatedChatMessage, err
}

func (connection *chatMessageRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.ChatMessage{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
