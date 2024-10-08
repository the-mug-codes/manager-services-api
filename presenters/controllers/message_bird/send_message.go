package message_bird

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	messageBird "github.com/the-mug-codes/service-manager-api/adapters/messagebird"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func SendMessage(context *gin.Context) {
	var newMessage *messageBird.NewMessage[entity.MessageContent]
	err := context.ShouldBindBodyWith(&newMessage, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
	response, err := messageBirdConnection.SendMessage(newMessage.To, newMessage.From, newMessage.Type, newMessage.Content, &newMessage.ReplyTo.ID)
	if err != nil {
		helper.ErrorResponse(context, 401, "cannot send message", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
