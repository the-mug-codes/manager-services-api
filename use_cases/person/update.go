package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(person entity.PersonRepository, dataToUpdate entity.Person) (updatedData *entity.Person, err error) {
	return person.Update(&dataToUpdate)
}
