package people

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(person entity.PersonRepository, page int, pageSize int) (readData *[]entity.Person, pagination *utils.Pagination, err error) {
	return person.ReadAll(page, pageSize)
}
