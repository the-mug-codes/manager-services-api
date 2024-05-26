package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func Update(invoice entity.InvoiceRepository, dataToInsert entity.Invoice) (updatedData *entity.Invoice, err error) {
	return invoice.Update(&dataToInsert)
}
