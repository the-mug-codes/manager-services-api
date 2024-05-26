package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(label entity.LabelRepository, id uuid.UUID) (readData *entity.Label, err error) {
	return label.Read(id)
}
