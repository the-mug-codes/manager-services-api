package message_bird

import (
	"github.com/gin-gonic/gin"
	messageBird "github.com/kodit-tecnologia/service-manager/adapters/messagebird"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

func GetAllContacts(context *gin.Context) {
	messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
	response, err := messageBirdConnection.GetAllContacts()
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot get contacts", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
