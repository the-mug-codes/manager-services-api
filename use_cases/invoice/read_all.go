package people

import (
	entity "github.com/kodit-tecnologia/service-manager/entities"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
)

func ReadAll(invoice entity.InvoiceRepository, page int, pageSize int) (readData *[]entity.Invoice, pagination *utils.Pagination, err error) {
	return invoice.ReadAll(page, pageSize)
}
