package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(productCategory entity.ProductCategoryRepository, dataToInsert entity.ProductCategory) (updatedData *entity.ProductCategory, err error) {
	return productCategory.Update(&dataToInsert)
}
