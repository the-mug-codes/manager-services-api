package message_bird

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	messageBird "github.com/the-mug-codes/service-manager-api/adapters/messagebird"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func GetContact(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot get contact", "id not provided")
		return
	}
	messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
	response, err := messageBirdConnection.GetContact(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot get contact", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
