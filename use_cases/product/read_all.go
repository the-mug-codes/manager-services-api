package people

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(product entity.ProductRepository, page int, pageSize int) (readData *[]entity.Product, pagination *utils.Pagination, err error) {
	return product.ReadAll(page, pageSize)
}
