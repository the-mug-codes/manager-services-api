package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(contract entity.ContractRepository, page int, pageSize int) (readData *[]entity.Contract, pagination *utils.Pagination, err error) {
	return contract.ReadAll(page, pageSize)
}
