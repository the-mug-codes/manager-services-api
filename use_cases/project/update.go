package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(project entity.ProjectRepository, dataToInsert entity.Project) (updatedData *entity.Project, err error) {
	return project.Update(&dataToInsert)
}
