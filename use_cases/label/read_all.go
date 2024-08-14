package people

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(label entity.LabelRepository, page int, pageSize int) (readData *[]entity.Label, pagination *utils.Pagination, err error) {
	return label.ReadAll(page, pageSize)
}
