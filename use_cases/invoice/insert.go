package people

import (
	"fmt"
	"strings"

	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func generateCode(data entity.Invoice) string {
	var code strings.Builder
	code.WriteString("INV-")
	initials := strings.Split(data.ContractID.String(), "-")[2]
	code.WriteString(initials)
	code.WriteString(fmt.Sprintf("-%s", data.DueDate.Format("0106")))
	return strings.ToUpper(code.String())
}

func Insert(invoice entity.InvoiceRepository, dataToInsert entity.Invoice) (insertedData *entity.Invoice, err error) {
	dataToInsert.Code = generateCode(dataToInsert)
	return invoice.Insert(&dataToInsert)
}
