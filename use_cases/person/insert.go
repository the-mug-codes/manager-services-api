package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(person entity.PersonRepository, dataToInsert entity.Person) (insertedData *entity.Person, err error) {
	return person.Insert(&dataToInsert)
}
