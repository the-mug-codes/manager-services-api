package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadByEmail(person entity.PersonRepository, email string) (readData *entity.Person, err error) {
	return person.ReadByEmail(email)
}
