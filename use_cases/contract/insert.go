package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Insert(contract entity.ContractRepository, dataToInsert entity.Contract) (insertedData *entity.Contract, err error) {
	return contract.Insert(&dataToInsert)
}
