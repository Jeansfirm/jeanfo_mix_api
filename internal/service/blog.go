package service

import (
	blog_definition "jeanfo_mix/internal/definition/blog"
	"jeanfo_mix/internal/model"

	"gorm.io/gorm"
)

type BlogService struct {
	DB *gorm.DB
}

func (s *BlogService) CreateArticle(article *model.Article) error {
	return s.DB.Create(article).Error
}

func (s *BlogService) ListArticle(req *blog_definition.ListArticleReq) ([]*model.Article, error) {
	query := s.DB
	if req.UserID != 0 {
		query = query.Where(&model.Article{UserID: int32(req.UserID)})
	}
	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	var articles []*model.Article
	err := query.Find(&articles).Error

	return articles, err
}

func (s *BlogService) CreateComment(comment *model.Comment) error {
	return s.DB.Create(comment).Error
}
