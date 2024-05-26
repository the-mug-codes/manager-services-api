package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(product entity.ProductRepository, dataToInsert entity.Product) (insertedData *entity.Product, err error) {
	return product.Insert(&dataToInsert)
}
