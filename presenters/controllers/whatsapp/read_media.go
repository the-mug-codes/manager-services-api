package whatsapp

import (
	"os"

	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	"github.com/the-mug-codes/service-manager-api/adapters/whatsapp"
)

// @Summary		Read a WhatsApp media file
// @Description	Reads a WhatsApp media file.
// @Tags			WhatsApp
// @Accept			json
// @Produce		json
// @Param			payload	body		whatsapp.SendInteractiveMessage	true	"payload"
// @Param			id				path		uuid.UUID	true	"ID"
// @Success		200		{file} application/octet-stream
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/whatsapp/media/:id [get]
func ReadMedia(context *gin.Context) {
	whatsappConnection := whatsapp.Connect(os.Getenv("WHATSAPP_ACCOUNT"))
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 404, "cannot read media", "id not provided")
		return
	}
	response, mimeType, err := whatsappConnection.ReadMedia(id)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read media", err.Error())
		return
	}
	context.Data(200, mimeType, response)
}
