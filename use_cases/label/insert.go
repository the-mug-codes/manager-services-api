package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(label entity.LabelRepository, dataToInsert entity.Label) (insertedData *entity.Label, err error) {
	return label.Insert(&dataToInsert)
}
