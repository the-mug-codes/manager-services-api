package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(project entity.ProjectRepository, id uuid.UUID) (success bool, err error) {
	err = project.Delete(id)
	if err != nil {
		return false, err
	}
	return true, err
}
