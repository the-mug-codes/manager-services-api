package people

import (
	"github.com/google/uuid"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Delete(contract entity.ContractRepository, id uuid.UUID) (err error) {
	return contract.Delete(id)
}
