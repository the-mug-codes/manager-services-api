package whatsapp

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	websocket "github.com/the-mug-codes/service-manager-api/adapters/websocket"
	whatsapp "github.com/the-mug-codes/service-manager-api/adapters/whatsapp"
	repositories "github.com/the-mug-codes/service-manager-api/repositories/chat"
	chat "github.com/the-mug-codes/service-manager-api/use_cases/chat/message"
)

func MessageReceived(context *gin.Context) {
	whatsappConnection := whatsapp.Connect(os.Getenv("WHATSAPP_ACCOUNT"))
	var requestBody *whatsapp.ReceiveMessage
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	websocketChat := websocket.GetWebsocketChat()
	err = whatsappConnection.ReceiveMessage(requestBody, func(message *whatsapp.MessageParsed) {
		go chat.SaveWhatsAppMessage(repositories.ChatSession(context), repositories.ChatMessage(context), websocketChat, message)
	})
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot receive message", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 201)
}
