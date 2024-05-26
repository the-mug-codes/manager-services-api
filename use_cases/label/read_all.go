package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(label entity.LabelRepository, page int, pageSize int) (readData *[]entity.Label, pagination *utils.Pagination, err error) {
	return label.ReadAll(page, pageSize)
}
