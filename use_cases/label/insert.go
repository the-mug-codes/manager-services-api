package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(label entity.LabelRepository, dataToInsert entity.Label) (insertedData *entity.Label, err error) {
	return label.Insert(&dataToInsert)
}
