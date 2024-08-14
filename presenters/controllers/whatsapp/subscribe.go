package whatsapp

import (
	"os"

	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

func Subscribe(context *gin.Context) {
	mode := context.Query("hub.mode")
	challenge := context.Query("hub.challenge")
	verifyToken := context.Query("hub.verify_token")
	if mode != "subscribe" && verifyToken != os.Getenv("WHATSAPP_VALIDATION") {
		helper.ErrorResponse(context, 400, "cannot subscribe webhook", "invalid mode or token")
		return
	}
	context.Data(200, "text/plain; charset=utf-8", []byte(challenge))
}
