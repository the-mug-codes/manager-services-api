package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(project entity.ProjectRepository, id uuid.UUID) (readData *entity.Project, err error) {
	return project.Read(id)
}
