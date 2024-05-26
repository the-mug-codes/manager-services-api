package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(person entity.PersonRepository, page int, pageSize int) (readData *[]entity.Person, pagination *utils.Pagination, err error) {
	return person.ReadAll(page, pageSize)
}
