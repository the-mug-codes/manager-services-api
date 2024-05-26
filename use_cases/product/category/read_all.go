package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(productCategory entity.ProductCategoryRepository, page int, pageSize int) (readData *[]entity.ProductCategory, pagination *utils.Pagination, err error) {
	return productCategory.ReadAll(page, pageSize)
}
