package controller

import (
	"database/sql"
	blog_definition "jeanfo_mix/internal/definition/blog"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/service"
	context_util "jeanfo_mix/util/context"
	request_util "jeanfo_mix/util/request"
	response_util "jeanfo_mix/util/response"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Service *service.BlogService
}

// @Summary Blog: Create Article
// @Tags Blog
// @Param article body blog_definition.CreateArticleReq true "create article"
// @Router /api/blog/articles [post]
// @Security BearerAuth
func (c *BlogController) CreateArticle(ctx *gin.Context) {
	req := request_util.NewRequest[blog_definition.CreateArticleReq](ctx).Data

	userID := context_util.NewHttpContext(ctx).SessionData().UserID
	article := model.Article{
		UserID: int32(userID), Ttile: req.Title, Content: req.Content, PlainText: req.PlainText,
	}
	err := c.Service.CreateArticle(&article)
	if err != nil {
		response_util.NewResponse(ctx).SetMsg("create articel fail: " + err.Error()).FailBadRequest()
		return
	}

	response_util.NewResponse(ctx).SetMsg("create article successfully").SetData(article).Success()
}

// @Summary Blog: List Article
// @Tags Blog
// @Param query query blog_definition.ListArticleReq true "list article"
// @Router /api/blog/articles [get]
// @Security BearerAuth
func (c *BlogController) ListArticle(ctx *gin.Context) {
	req := request_util.NewRequest[blog_definition.ListArticleReq](ctx)
	reqData := req.Data
	reqData.AutoFill()
	articles, err := c.Service.ListArticle(reqData)

	if err != nil {
		response_util.NewResponse(ctx).SetMsg("list articles fail: " + err.Error()).FailBadRequest()
		return
	}

	response_util.NewResponse(ctx).SetData(articles).Success()
}

// @Summary Blog: List Article My
// @Tags Blog
// @Param query query blog_definition.ListArticleReq true "list article my"
// @Router /api/blog/articles/my [get]
// @Security BearerAuth
func (c *BlogController) ListArticleMy(ctx *gin.Context) {
	sessData := context_util.NewHttpContext(ctx).SessionData()
	req := request_util.NewRequest[blog_definition.ListArticleReq](ctx)
	reqData := req.Data
	reqData.AutoFill()
	reqData.UserID = sessData.UserID
	articles, err := c.Service.ListArticle(reqData)

	if err != nil {
		response_util.NewResponse(ctx).SetMsg("list my articles fail: " + err.Error()).FailBadRequest()
		return
	}

	response_util.NewResponse(ctx).SetData(articles).Success()
}

// @Summary Blog: Create Comment
// @Tags Blog
// @Param comment body blog_definition.CreateCommentReq true "create comment"
// @Router /api/blog/comments [post]
// @Security BearerAuth
func (c *BlogController) CreateComment(ctx *gin.Context) {
	req := request_util.NewRequest[blog_definition.CreateCommentReq](ctx).Data

	userID := context_util.NewHttpContext(ctx).SessionData().UserID
	comment := model.Comment{
		UserID: int32(userID), ArticleID: req.ArticleID,
		Content: req.Content, PlainText: req.PlainText,
	}
	if req.CommentID == nil {
		comment.CommentID = sql.NullInt32{Valid: false}
	} else {
		comment.CommentID = sql.NullInt32{Int32: int32(*req.CommentID), Valid: true}
	}

	err := c.Service.CreateComment(&comment)
	if err != nil {
		response_util.NewResponse(ctx).SetMsg("create comment fail: " + err.Error()).FailBadRequest()
		return
	}

	response_util.NewResponse(ctx).SetMsg("create comment successfully").SetData(comment).Success()
}
