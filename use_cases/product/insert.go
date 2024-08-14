package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(product entity.ProductRepository, dataToInsert entity.Product) (insertedData *entity.Product, err error) {
	return product.Insert(&dataToInsert)
}
