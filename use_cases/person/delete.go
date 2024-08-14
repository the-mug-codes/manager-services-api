package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(person entity.PersonRepository, id uuid.UUID) (err error) {
	return person.Delete(id)
}
