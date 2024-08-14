package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(contract entity.ContractRepository, dataToInsert entity.Contract) (updatedData *entity.Contract, err error) {
	return contract.Update(&dataToInsert)
}
