package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(product entity.ProductRepository, id uuid.UUID) (readData *entity.Product, err error) {
	return product.Read(id)
}
