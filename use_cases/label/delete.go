package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(label entity.LabelRepository, id uuid.UUID) (err error) {
	return label.Delete(id)
}
