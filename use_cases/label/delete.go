package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Delete(label entity.LabelRepository, id uuid.UUID) (err error) {
	return label.Delete(id)
}
