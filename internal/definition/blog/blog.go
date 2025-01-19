package blog_definition

import (
	"jeanfo_mix/internal/definition"
)

type CreateArticleReq struct {
	definition.BaseReq

	Title     string `json:"Title" binding:"required"`
	Content   string `json:"Content" binding:"required"`
	PlainText string `json:"PlainText"`
}

type ListArticleReq struct {
	definition.PageReq

	UserID int `json:"UserID"`
}

type CreateCommentReq struct {
	definition.BaseReq

	Content   string `json:"Content" binding:"required"`
	PlainText string `json:"PlainText"`
	ArticleID int    `json:"ArticleID" binding:"required"`
	CommentID *int   `json:"CommentID"`

	UserID int `json:"UserID"`
}
