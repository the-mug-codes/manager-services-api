package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Insert(contract entity.ContractRepository, dataToInsert entity.Contract) (insertedData *entity.Contract, err error) {
	return contract.Insert(&dataToInsert)
}
