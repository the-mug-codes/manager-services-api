package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(contract entity.ContractRepository, id uuid.UUID) (readData *entity.Contract, err error) {
	return contract.Read(id)
}
