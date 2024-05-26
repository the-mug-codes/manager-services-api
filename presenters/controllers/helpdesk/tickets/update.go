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

type updateHelpDeskTicket struct {
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

func (data updateHelpDeskTicket) dataToUpdate(id uuid.UUID) (dataToUpdate *entity.HelpDeskTicket) {
	return &entity.HelpDeskTicket{
		ID:        id,
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
}

// @Summary		Change a HelpDesk Ticket
// @Description	Updates a helpdesk ticket by idin database.
// @Tags			HelpDesk - Tickets
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID				true	"ID"
// @Param			payload	body		updateHelpDeskTicket	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entities.HelpDeskTicket]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/helpdesk/tickets/{id} [put]
func Update(context *gin.Context) {
	var requestBody *updateHelpDeskTicket
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
	response, err := helpDeskTicket.Update(repository.HelpDeskTicket(context), *requestBody.dataToUpdate(uuid))
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
