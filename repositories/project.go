package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	database "github.com/the-mug-codes/adapters-service-api/database"
	utils "github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type projectRepository struct {
	context  *gin.Context
	database *gorm.DB
}

func Project(context *gin.Context) entity.ProjectRepository {
	return &projectRepository{
		context:  context,
		database: database.GetConnection(),
	}
}

func (connection *projectRepository) Insert(projectData *entity.Project) (insertedProject *entity.Project, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &projectData.PersonID) != nil {
		return insertedProject, err
	}
	err = connection.database.Omit("Person").Create(&projectData).Error
	if err != nil {
		return insertedProject, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Person").First(&insertedProject, projectData.ID).Error
	if err != nil {
		return insertedProject, err
	}
	return insertedProject, err
}

func (connection *projectRepository) ReadAll(page int, pageSize int) (projects *[]entity.Project, pagination *utils.Pagination, err error) {
	var totalRegisters int64
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Find(&entity.Project{}).Count(&totalRegisters).Error
	if err != nil {
		return projects, pagination, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Scopes(utils.Paginate(page, pageSize, totalRegisters, pagination)).Preload("Labels").Preload("Person").Find(&projects).Error
	if err != nil {
		return projects, pagination, err
	}
	return projects, pagination, err
}

func (connection *projectRepository) Read(id uuid.UUID) (project *entity.Project, err error) {
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Person").First(&project, id).Error
	if err != nil {
		return project, err
	}
	return project, err
}

func (connection *projectRepository) Update(projectData *entity.Project) (updatedProject *entity.Project, err error) {
	if utils.VerifyOwnerRestriction(connection.context, &projectData.PersonID) != nil {
		return updatedProject, err
	}
	err = connection.database.Omit("CreatedAt").Save(&projectData).Error
	if err != nil {
		return updatedProject, err
	}
	err = connection.database.Scopes(utils.FilterDataOwnerRestriction(connection.context, "id")).Preload("Labels").Preload("Person").First(&updatedProject, projectData.ID).Error
	if err != nil {
		return updatedProject, err
	}
	return updatedProject, err
}

func (connection *projectRepository) Delete(id uuid.UUID) (err error) {
	err = connection.database.Delete(&entity.Project{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
