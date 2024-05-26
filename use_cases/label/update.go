package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(label entity.LabelRepository, dataToInsert entity.Label) (updatedData *entity.Label, err error) {
	return label.Update(&dataToInsert)
}
