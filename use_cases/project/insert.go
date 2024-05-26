package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(project entity.ProjectRepository, dataToInsert entity.Project) (insertedData *entity.Project, err error) {
	return project.Insert(&dataToInsert)
}
