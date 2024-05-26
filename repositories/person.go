package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type personRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func Person(context *gin.Context) entity.PersonRepository {
	return &personRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *personRepository) Insert(personData *entity.Person) (insertedPerson *entity.Person, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &personData.ID) != nil {
		return insertedPerson, err
	}
	err = connection.database.Create(&personData).Error
	if err != nil {
		return insertedPerson, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Addresses").Preload("Emails").Preload("Phones").First(&insertedPerson, personData.ID).Error
	if err != nil {
		return insertedPerson, err
	}
	return insertedPerson, err
}

func (connection *personRepository) ReadAll(page int, pageSize int) (person *[]entity.Person, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Find(&entity.Person{}).Count(&totalRegisters).Error
	if err != nil {
		return person, pagination, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Preload("Addresses").Preload("Emails").Preload("Phones").Find(&person).Error
	if err != nil {
		return person, pagination, err
	}
	return person, pagination, err
}

func (connection *personRepository) Read(id uuid.UUID) (person *entity.Person, err error) {
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Addresses").Preload("Emails").Preload("Phones").First(&person, id).Error
	if err != nil {
		return person, err
	}
	return person, err
}

func (connection *personRepository) ReadByEmail(email string) (person *entity.Person, err error) {
	var personEmail *entity.PersonEmail
	err = connection.database.Where("email = ?", email).First(&personEmail).Error
	if err != nil {
		return person, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Addresses").Preload("Emails").Preload("Phones").First(&person, personEmail.PersonID).Error
	if err != nil {
		return person, err
	}
	return person, err
}

func (connection *personRepository) ReadByPhoneNumber(countryCode int, areaCode int, phoneNumber int) (person *entity.Person, err error) {
	var personPhone *entity.PersonPhone
	err = connection.database.Where("country_code = ? AND area_code = ? AND phone_number = ?", countryCode, areaCode, phoneNumber).First(&personPhone).Error
	if err != nil {
		return person, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Addresses").Preload("Emails").Preload("Phones").First(&person, personPhone.PersonID).Error
	if err != nil {
		return person, err
	}
	return person, err
}

func (connection *personRepository) Update(personData *entity.Person) (updatedPersons *entity.Person, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &personData.ID) != nil {
		return updatedPersons, err
	}
	err = connection.database.Omit("CreatedAt").Save(&personData).Error
	if err != nil {
		return updatedPersons, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Addresses").Preload("Emails").Preload("Phones").First(&updatedPersons, personData.ID).Error
	if err != nil {
		return updatedPersons, err
	}
	return updatedPersons, err
}

func (connection *personRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.Person{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
