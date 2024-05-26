package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(person entity.PersonRepository, dataToInsert entity.Person) (insertedData *entity.Person, err error) {
	return person.Insert(&dataToInsert)
}
