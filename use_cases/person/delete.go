package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Delete(person entity.PersonRepository, id uuid.UUID) (err error) {
	return person.Delete(id)
}
