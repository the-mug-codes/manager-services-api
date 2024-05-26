package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Delete(contract entity.ContractRepository, id uuid.UUID) (err error) {
	return contract.Delete(id)
}
