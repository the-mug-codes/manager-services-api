package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Delete(product entity.ProductRepository, id uuid.UUID) (err error) {
	return product.Delete(id)
}
