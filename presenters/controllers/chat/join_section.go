package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	websocket "github.com/the-mug-codes/service-manager-api/adapters/websocket"
	repositories "github.com/the-mug-codes/service-manager-api/repositories/chat"
	chat "github.com/the-mug-codes/service-manager-api/use_cases/chat/websocket"
)

func JoinSection(context *gin.Context) {
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot join chat session", "id not provided")
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "invalid uuid", err.Error())
		return
	}
	websocketConnection, err := websocket.GetWebSocketUpgrader(context)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot join chat session", err.Error())
		return
	}
	websocketChat := websocket.GetWebsocketChat()
	err = chat.JoinChatSection(repositories.ChatMessage(context), websocketChat, websocketConnection, uuid)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot join chat session", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 200)
}
