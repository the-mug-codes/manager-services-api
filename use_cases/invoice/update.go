package people

import (
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func Update(invoice entity.InvoiceRepository, dataToInsert entity.Invoice) (updatedData *entity.Invoice, err error) {
	return invoice.Update(&dataToInsert)
}
