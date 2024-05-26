package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Delete(productCategory entity.ProductCategoryRepository, id uuid.UUID) (err error) {
	return productCategory.Delete(id)
}
