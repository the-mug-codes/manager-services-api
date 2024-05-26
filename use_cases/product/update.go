package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(product entity.ProductRepository, dataToInsert entity.Product) (updatedData *entity.Product, err error) {
	return product.Update(&dataToInsert)
}
