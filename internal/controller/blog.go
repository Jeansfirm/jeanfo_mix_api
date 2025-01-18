package controller

import (
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/service"
	context_util "jeanfo_mix/util/context"
	response_util "jeanfo_mix/util/response"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Service *service.BlogService
}

type CreateArticleReq struct {
	Title     string `binding:"required"`
	Content   string `binding:"required"`
	PlainText string
}

// @Summary Blog: Create Article
// @Tags Blog
// @Param article body CreateArticleReq true "create article"
// @Router /api/blog/articles [post]
// @Security BearerAuth
func (c *BlogController) CreateArticle(ctx *gin.Context) {
	var req CreateArticleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response_util.NewResponse(ctx).SetMsg("params error: " + err.Error()).FailBadRequest()
		return
	}

	userID := context_util.NewHttpContext(ctx).SessionData().UserID
	article := model.Article{
		UserID: userID, Ttile: req.Title, Content: req.Content, PlainText: req.PlainText,
	}
	err := c.Service.CreateArticle(&article)
	if err != nil {
		response_util.NewResponse(ctx).SetMsg("create articel fail: " + err.Error()).FailBadRequest()
		return
	}

	response_util.NewResponse(ctx).SetMsg("create article successfully").SetData(article).Success()
}
