package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(project entity.ProjectRepository, dataToInsert entity.Project) (updatedData *entity.Project, err error) {
	return project.Update(&dataToInsert)
}
