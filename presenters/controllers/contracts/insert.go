package attachment

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	contract "github.com/kodit-tecnologia/service-manager/use_cases/contract"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type insertContract struct {
	ContractModel     string       `json:"contract_model" binding:"required"`
	Labels            []*uuid.UUID `json:"labels,omitempty"`
	PersonID          uuid.UUID    `json:"person_id"`
	ProjectID         uuid.UUID    `json:"project_id"`
	Description       string       `json:"description" binding:"required"`
	Comments          string       `json:"comments" binding:"required"`
	AmountValue       float64      `json:"amount_value" binding:"required"`
	DiscountValue     float64      `json:"discount_value" binding:"required"`
	ExtraAmountValue  float64      `json:"extra_amount_value" binding:"required"`
	TotalAmountValue  float64      `json:"total_amount_value" binding:"required"`
	InstallmentNumber *int         `json:"installment_number,omitempty"`
	DueDate           *time.Time   `json:"due_date" binding:"required"`
	SignedAt          *time.Time   `json:"signed_at,omitempty"`
	Document          string       `json:"document,omitempty"`
	ResellerID        *uuid.UUID   `json:"reseller_id"`
	Products          []*uuid.UUID `json:"products,omitempty"`
}

func (data insertContract) dataToInsert() (dataToInsert *entity.Contract) {
	dataToInsert = &entity.Contract{
		ContractModel:     dataToInsert.ContractModel,
		PersonID:          dataToInsert.PersonID,
		ProjectID:         dataToInsert.ProjectID,
		Description:       dataToInsert.Description,
		Comments:          dataToInsert.Comments,
		AmountValue:       dataToInsert.AmountValue,
		DiscountValue:     dataToInsert.DiscountValue,
		ExtraAmountValue:  dataToInsert.ExtraAmountValue,
		TotalAmountValue:  dataToInsert.TotalAmountValue,
		InstallmentNumber: dataToInsert.InstallmentNumber,
		DueDate:           dataToInsert.DueDate,
		SignedAt:          dataToInsert.SignedAt,
		Document:          dataToInsert.Document,
		ResellerID:        dataToInsert.ResellerID,
	}
	for _, label := range data.Labels {
		dataToInsert.Labels = append(dataToInsert.Labels, &entity.Label{
			ID: *label,
		})
	}
	for _, product := range data.Products {
		dataToInsert.Products = append(dataToInsert.Products, &entity.Product{
			ID: *product,
		})
	}
	return dataToInsert
}

// @Summary		Create a Contract
// @Description	Creates a new contract in database.
// @Tags			Contracts
// @Accept			json
// @Produce		json
// @Param			payload	body		insertContract	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Contract]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/contracts [post]
func Insert(context *gin.Context) {
	var requestBody *insertContract
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := contract.Insert(repository.Contract(context), *requestBody.dataToInsert())
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
