package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	"gorm.io/gorm"
)

type chatSessionRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func ChatSession(context *gin.Context) entity.ChatSessionRepository {
	return &chatSessionRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *chatSessionRepository) Insert(chatSessionData *entity.ChatSession) (insertedChatSession *entity.ChatSession, err error) {
	err = connection.database.Create(&chatSessionData).Error
	if err != nil {
		return insertedChatSession, err
	}
	err = connection.database.First(&insertedChatSession, chatSessionData.ID).Error
	if err != nil {
		return insertedChatSession, err
	}
	return insertedChatSession, err
}

func (connection *chatSessionRepository) ReadAll(page int, pageSize int) (chatSession *[]entity.ChatSession, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Find(&entity.ChatSession{}).Count(&totalRegisters).Error
	if err != nil {
		return chatSession, pagination, err
	}
	err = connection.database.Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Find(&chatSession).Error
	if err != nil {
		return chatSession, pagination, err
	}
	return chatSession, pagination, err
}

func (connection *chatSessionRepository) Read(id uuid.UUID) (chatSession *entity.ChatSession, err error) {
	err = connection.database.First(&chatSession, id).Error
	if err != nil {
		return chatSession, err
	}
	return chatSession, err
}

func (connection *chatSessionRepository) ReadByPhoneNumber(phoneNumber int) (chatSession *entity.ChatSession, err error) {
	err = connection.database.Where("phone_number = ?", phoneNumber).First(&chatSession).Error
	if err != nil {
		return chatSession, err
	}
	return chatSession, err
}

func (connection *chatSessionRepository) Update(chatSessionData *entity.ChatSession) (updatedChatSession *entity.ChatSession, err error) {
	err = connection.database.First(&updatedChatSession, chatSessionData.ID).Error
	if err != nil {
		return updatedChatSession, err
	}
	return updatedChatSession, err
}

func (connection *chatSessionRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.ChatSession{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
