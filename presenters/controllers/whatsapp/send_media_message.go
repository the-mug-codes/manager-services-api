package whatsapp

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	whatsapp "github.com/the-mug-codes/service-manager-api/adapters/whatsapp"
)

// @Summary		Send a WhatsApp media message
// @Description	Sends a new WhatsApp media message.
// @Tags			WhatsApp
// @Accept			json
// @Produce		json
// @Param			payload	body		whatsapp.SendMediaMessage	true	"payload"
// @Param			phone				path		uuid.UUID	true	"Phone"
// @Success		201		{object}	helper.ResponseNone
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/whatsapp/messages/:phone/media [post]
func SendMediaMessage(context *gin.Context) {
	whatsappConnection := whatsapp.Connect(os.Getenv("WHATSAPP_ACCOUNT"))
	phone, havePhone := context.Params.Get("phone")
	if !havePhone {
		helper.ErrorResponse(context, 400, "cannot send message", "id not provided")
		return
	}
	mediaType, haveMediaType := context.GetQuery("type")
	if !haveMediaType {
		helper.ErrorResponse(context, 400, "cannot send message", "message type not provided")
		return
	}
	var dataToSend *whatsapp.SendMediaMessage
	err := context.ShouldBindBodyWith(&dataToSend, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	err = whatsappConnection.SendMediaMessage(phone, whatsapp.MediaType(mediaType), *dataToSend)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot send message", err.Error())
		return
	}
	helper.SuccessResponseNone(context, 201)
}
