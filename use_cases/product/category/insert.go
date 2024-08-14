package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(productCategory entity.ProductCategoryRepository, dataToInsert entity.ProductCategory) (insertedData *entity.ProductCategory, err error) {
	return productCategory.Insert(&dataToInsert)
}
