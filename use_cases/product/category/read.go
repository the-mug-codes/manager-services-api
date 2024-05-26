package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(productCategory entity.ProductCategoryRepository, id uuid.UUID) (readData *entity.ProductCategory, err error) {
	return productCategory.Read(id)
}
