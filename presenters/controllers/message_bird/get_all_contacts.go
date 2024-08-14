package message_bird

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	messageBird "github.com/the-mug-codes/service-manager-api/adapters/messagebird"
	entity "github.com/the-mug-codes/service-manager-api/entities"
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
