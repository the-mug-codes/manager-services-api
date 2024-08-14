package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(productCategory entity.ProductCategoryRepository, id uuid.UUID) (err error) {
	return productCategory.Delete(id)
}
