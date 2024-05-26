package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(contract entity.ContractRepository, dataToInsert entity.Contract) (updatedData *entity.Contract, err error) {
	return contract.Update(&dataToInsert)
}
