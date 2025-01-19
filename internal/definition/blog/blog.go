package blog_definition

import (
	"jeanfo_mix/internal/definition"
)

type CreateArticleReq struct {
	definition.BaseReq

	Title     string `binding:"required"`
	Content   string `binding:"required"`
	PlainText string
}

type ListArticleReq struct {
	definition.PageReq

	UserID int `json:"UserID"`
}
