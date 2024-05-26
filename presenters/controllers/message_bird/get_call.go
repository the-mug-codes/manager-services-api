package message_bird

import (
	"github.com/gin-gonic/gin"
	messageBird "github.com/kodit-tecnologia/service-manager/adapters/messagebird"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

func GetCall(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot get call", "id not provided")
		return
	}
	messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
	response, err := messageBirdConnection.GetCall(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot get call", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
