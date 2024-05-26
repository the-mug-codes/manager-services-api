package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(project entity.ProjectRepository, id uuid.UUID) (readData *entity.Project, err error) {
	return project.Read(id)
}
