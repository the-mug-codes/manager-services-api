package people

import (
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func ReadAll(invoice entity.InvoiceRepository, page int, pageSize int) (readData *[]entity.Invoice, pagination *utils.Pagination, err error) {
	return invoice.ReadAll(page, pageSize)
}
