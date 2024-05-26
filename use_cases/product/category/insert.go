package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(productCategory entity.ProductCategoryRepository, dataToInsert entity.ProductCategory) (insertedData *entity.ProductCategory, err error) {
	return productCategory.Insert(&dataToInsert)
}
