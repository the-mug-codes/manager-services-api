package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(person entity.PersonRepository, id uuid.UUID) (readData *entity.Person, err error) {
	return person.Read(id)
}
