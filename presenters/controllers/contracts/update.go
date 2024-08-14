package attachment

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	contract "github.com/the-mug-codes/service-manager-api/use_cases/contract"
)

type updateContract struct {
	ContractModel string `json:"contract_model" binding:"required"`
	Labels        []*struct {
		ID uuid.UUID `json:"id"`
	} `json:"labels,omitempty"`
	PersonID          uuid.UUID  `json:"person_id"`
	ProjectID         uuid.UUID  `json:"project_id"`
	Description       string     `json:"description" binding:"required"`
	Comments          string     `json:"comments" binding:"required"`
	AmountValue       float64    `json:"amount_value" binding:"required"`
	DiscountValue     float64    `json:"discount_value" binding:"required"`
	ExtraAmountValue  float64    `json:"extra_amount_value" binding:"required"`
	TotalAmountValue  float64    `json:"total_amount_value" binding:"required"`
	InstallmentNumber *int       `json:"installment_number,omitempty"`
	DueDate           *time.Time `json:"due_date" binding:"required"`
	SignedAt          *time.Time `json:"signed_at,omitempty"`
	Document          string     `json:"document,omitempty"`
	ResellerID        *uuid.UUID `json:"reseller_id"`
	Products          []*struct {
		ID uuid.UUID `json:"id"`
	} `json:"products,omitempty"`
}

func (data updateContract) dataToUpdate(id uuid.UUID) (dataToUpdate *entity.Contract) {
	dataToUpdate = &entity.Contract{
		ID:                id,
		ContractModel:     data.ContractModel,
		PersonID:          data.PersonID,
		ProjectID:         data.ProjectID,
		Description:       data.Description,
		Comments:          data.Comments,
		AmountValue:       data.AmountValue,
		DiscountValue:     data.DiscountValue,
		ExtraAmountValue:  data.ExtraAmountValue,
		TotalAmountValue:  data.TotalAmountValue,
		InstallmentNumber: data.InstallmentNumber,
		DueDate:           data.DueDate,
		SignedAt:          data.SignedAt,
		Document:          data.Document,
		ResellerID:        data.ResellerID,
	}
	for _, label := range data.Labels {
		dataToUpdate.Labels = append(dataToUpdate.Labels, &entity.Label{
			ID: label.ID,
		})
	}
	for _, product := range data.Products {
		dataToUpdate.Products = append(dataToUpdate.Products, &entity.Product{
			ID: product.ID,
		})
	}
	return dataToUpdate
}

// @Summary		Change a Contract
// @Description	Updates a ccontract by idin database.
// @Tags			Contracts
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID		true	"ID"
// @Param			payload	body		updateContract	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entities.Contract]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/contracts/{id} [put]
func Update(context *gin.Context) {
	var requestBody *updateContract
	id, haveId := context.Params.Get("id")
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil || !haveId {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	response, err := contract.Update(repository.Contract(context), *requestBody.dataToUpdate(uuid))
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
