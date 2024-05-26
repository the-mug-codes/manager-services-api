package attachment

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories/helpdesk"
	helpDeskTicket "github.com/kodit-tecnologia/service-manager/use_cases/helpdesk/ticket"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type insertHelpDeskTicket struct {
	Title     string       `json:"title" binding:"required"`
	Level     string       `json:"level" binding:"required"`
	Code      int          `json:"code" binding:"required"`
	Labels    []*uuid.UUID `json:"labels,omitempty"`
	PersonID  uuid.UUID    `json:"person_id"`
	ProjectID *uuid.UUID   `json:"project_id,omitempty"`
	DueDate   time.Time    `json:"due_date" binding:"required"`
	Source    string       `json:"source" binding:"required"`
	SourceID  string       `json:"source_id" binding:"required"`
	Status    *bool        `json:"status"`
}

func (data insertHelpDeskTicket) dataToInsert() (dataToInsert *entity.HelpDeskTicket) {
	dataToInsert = &entity.HelpDeskTicket{
		Title:     data.Title,
		Level:     data.Level,
		Code:      data.Code,
		PersonID:  data.PersonID,
		ProjectID: data.ProjectID,
		DueDate:   data.DueDate,
		Source:    data.Source,
		SourceID:  data.SourceID,
		Status:    *data.Status,
	}
	return dataToInsert
}

// @Summary		Create a HelpDesk Ticket
// @Description	Creates a new helpdesk ticket in database.
// @Tags			HelpDesk - Tickets
// @Accept			json
// @Produce		json
// @Param			payload	body		insertHelpDeskTicket	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.HelpDeskTicket]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/helpdesk/tickets [post]
func Insert(context *gin.Context) {
	var requestBody *insertHelpDeskTicket
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := helpDeskTicket.Insert(repository.HelpDeskTicket(context), *requestBody.dataToInsert())
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
