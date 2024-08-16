package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	websocket "github.com/the-mug-codes/service-manager-api/adapters/websocket"
	repositories "github.com/the-mug-codes/service-manager-api/repositories/chat"
	chat "github.com/the-mug-codes/service-manager-api/use_cases/chat/websocket"
)

type createSession struct {
	Email       *string `json:"email,omitempty"`
	PhoneNumber *int    `json:"phone_number,omitempty"`
}

func CreateSection(context *gin.Context) {
	var requestBody *createSession
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	websocketChat := websocket.GetWebsocketChat()
	sessionID, err := chat.CreateChatSection(repositories.ChatSession(context), websocketChat, requestBody.Email, requestBody.PhoneNumber)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot create chat session", err.Error())
		return
	}
	helper.SuccessResponseOne[uuid.UUID](context, 201, (*uuid.UUID)(&sessionID))
}
