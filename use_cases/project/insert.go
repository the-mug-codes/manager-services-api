package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(project entity.ProjectRepository, dataToInsert entity.Project) (insertedData *entity.Project, err error) {
	return project.Insert(&dataToInsert)
}
