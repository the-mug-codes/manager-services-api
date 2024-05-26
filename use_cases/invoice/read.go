package people

import (
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Read(invoice entity.InvoiceRepository, id uuid.UUID) (readData *entity.Invoice, err error) {
	return invoice.Read(id)
}
