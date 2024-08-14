package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(product entity.ProductRepository, id uuid.UUID) (readData *entity.Product, err error) {
	return product.Read(id)
}
