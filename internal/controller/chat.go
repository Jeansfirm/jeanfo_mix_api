package controller

import (
	"jeanfo_mix/internal/definition"
	chat_definition "jeanfo_mix/internal/definition/chat"
	chat_service "jeanfo_mix/internal/service/chat"
	context_util "jeanfo_mix/util/context"
	request_util "jeanfo_mix/util/request"
	reponse_util "jeanfo_mix/util/response"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	Service *chat_service.ChatService
}

func (c *ChatController) CreateConversion(ctx *gin.Context) {
	httpContext := context_util.NewHttpContext(ctx)
	userID := httpContext.SessionData().UserID

	req := request_util.NewRequest[chat_definition.CreateConversationReq](ctx).Data
	req.UserID = uint(userID)
	req.AutoFill()

	conv, err := c.Service.CreateConversation(req)
	if err != nil {
		reponse_util.New(ctx).SetMsg("create conversation fail: " + err.Error()).FailBadRequest()
	}

	reponse_util.New(ctx).SetData(conv).Success()
}

func (c *ChatController) ListConversation(ctx *gin.Context) {
	httpContext := context_util.NewHttpContext(ctx)
	userId := httpContext.SessionData().UserID

	req := request_util.NewRequest[definition.PageUserReq](ctx).Data
	req.UserID = uint(userId)

	total, convs, err := c.Service.ListConversation(req)
	if err != nil {
		reponse_util.New(ctx).SetMsg("list conversation fail: " + err.Error()).FailBadRequest()
		return
	}

	reponse_util.New(ctx).SetDataPaginated(total, convs, req.Page, req.PageSize).Success()
}

func (c *ChatController) ListMessage(ctx *gin.Context) {
	httpContext := context_util.NewHttpContext(ctx)
	userId := httpContext.SessionData().UserID

	req := request_util.NewRequest[chat_definition.ListMessageReq](ctx).Data
	req.UserID = uint(userId)

	total, msgs, err := c.Service.ListMessage(req)
	if err != nil {
		reponse_util.New(ctx).SetMsg("list message fail: " + err.Error()).FailBadRequest()
		return
	}

	reponse_util.New(ctx).SetDataPaginated(total, msgs, req.Page, req.PageSize).Success()
}
