package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(project entity.ProjectRepository, page int, pageSize int) (readData *[]entity.Project, pagination *utils.Pagination, err error) {
	return project.ReadAll(page, pageSize)
}
