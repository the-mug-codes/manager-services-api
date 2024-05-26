package message_bird

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	messageBird "github.com/kodit-tecnologia/service-manager/adapters/messagebird"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

func GetCallRecording(context *gin.Context) {
	messageBirdConnection := messageBird.Connect[entity.MessageContent, entity.NewMessageCreated]()
	fileType, isFile := context.GetQuery("file")
	_, isTranscription := context.GetQuery("transcription")
	id, haveId := context.Params.Get("id")
	if !haveId {
		helper.ErrorResponse(context, 400, "cannot get call recording", "id not provided")
		return
	}
	if isFile {
		file, err := messageBirdConnection.GetCallRecordingFile(id, fileType)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot get call recording transcription", err.Error())
			return
		}
		mimeType := fmt.Sprintf("audio/%s", fileType)
		content := fmt.Sprintf("attachment; filename=%s.%s", id, fileType)
		context.Header("Content-Disposition", content)
		context.Header("content-type", mimeType)
		context.Data(http.StatusOK, mimeType, *file)
		return
	}
	if isTranscription {
		transcription, err := messageBirdConnection.GetCallRecordingTranscription(id)
		if err != nil {
			helper.ErrorResponse(context, 400, "cannot get call recording file", err.Error())
			return
		}
		helper.SuccessResponseOne(context, 200, transcription)
		return
	}
	response, err := messageBirdConnection.GetCallRecording(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot get call recording", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
