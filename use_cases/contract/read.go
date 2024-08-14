package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Read(contract entity.ContractRepository, id uuid.UUID) (readData *entity.Contract, err error) {
	return contract.Read(id)
}
