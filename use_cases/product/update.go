package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(product entity.ProductRepository, dataToInsert entity.Product) (updatedData *entity.Product, err error) {
	return product.Update(&dataToInsert)
}
