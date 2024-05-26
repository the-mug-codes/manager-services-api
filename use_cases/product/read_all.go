package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(product entity.ProductRepository, page int, pageSize int) (readData *[]entity.Product, pagination *utils.Pagination, err error) {
	return product.ReadAll(page, pageSize)
}
