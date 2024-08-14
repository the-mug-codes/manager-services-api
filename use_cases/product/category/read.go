package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(productCategory entity.ProductCategoryRepository, id uuid.UUID) (readData *entity.ProductCategory, err error) {
	return productCategory.Read(id)
}
