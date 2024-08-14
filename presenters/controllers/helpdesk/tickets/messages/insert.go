package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories/helpdesk"
	helpDeskTicketMessage "github.com/the-mug-codes/service-manager-api/use_cases/helpdesk/ticket/message"
)

type insertTicketMessage struct {
	Text string `json:"description" binding:"required"`
}

func (data insertTicketMessage) dataToInsert(ticketID uuid.UUID) (dataToInsert *entity.HelpDeskTicketMessage) {
	return &entity.HelpDeskTicketMessage{
		TicketID: ticketID,
		Text:     data.Text,
	}
}

// @Summary		Create a HelpDesk Ticket Message
// @Description	Creates a new helpdesk ticket message in database.
// @Tags			HelpDesk - Tickets Messages
// @Accept			json
// @Produce		json
// @Param			payload	body		insertTicketMessage	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.HelpDeskTicketMessage]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/helpdesk/tickets/messages [post]
func Insert(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 404, "cannot read", "ticket id not provided")
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	var requestBody *insertTicketMessage
	err = context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := helpDeskTicketMessage.Insert(repository.HelpDeskTicketMessage(context), *requestBody.dataToInsert(uuid))
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
