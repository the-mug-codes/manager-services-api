package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func ReadByEmail(person entity.PersonRepository, email string) (readData *entity.Person, err error) {
	return person.ReadByEmail(email)
}
