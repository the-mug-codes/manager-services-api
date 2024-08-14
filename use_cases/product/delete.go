package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(product entity.ProductRepository, id uuid.UUID) (err error) {
	return product.Delete(id)
}
