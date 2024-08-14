package people

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(productCategory entity.ProductCategoryRepository, page int, pageSize int) (readData *[]entity.ProductCategory, pagination *utils.Pagination, err error) {
	return productCategory.ReadAll(page, pageSize)
}
