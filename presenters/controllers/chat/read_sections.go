package chat

import (
	"github.com/gin-gonic/gin"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	repositories "github.com/the-mug-codes/service-manager-api/repositories/chat"
	chat "github.com/the-mug-codes/service-manager-api/use_cases/chat/session"
)

func ReadSections(context *gin.Context) {
	page := helper.GetPageNumber(context)
	pageSize := helper.GetPageSize(context)
	response, pagination, err := chat.ReadAll(repositories.ChatSession(context), page, pageSize)
	if err != nil {
		helper.ErrorResponse(context, 404, "cannot read", err.Error())
		return
	}
	helper.SuccessResponseMany(context, 200, response, pagination)
}
