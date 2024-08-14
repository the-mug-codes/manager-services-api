package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(label entity.LabelRepository, dataToInsert entity.Label) (updatedData *entity.Label, err error) {
	return label.Update(&dataToInsert)
}
