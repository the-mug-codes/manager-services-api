package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(person entity.PersonRepository, dataToUpdate entity.Person) (updatedData *entity.Person, err error) {
	return person.Update(&dataToUpdate)
}
