package message_bird

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	messageBird "github.com/the-mug-codes/service-manager-api/adapters/messagebird"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func GetConversation(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot get conversation", "id not provided")
		return
	}
	messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
	response, err := messageBirdConnection.GetConversation(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot get conversation", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
