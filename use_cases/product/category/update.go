package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(productCategory entity.ProductCategoryRepository, dataToInsert entity.ProductCategory) (updatedData *entity.ProductCategory, err error) {
	return productCategory.Update(&dataToInsert)
}
