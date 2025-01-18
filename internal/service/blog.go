package service

import (
	"jeanfo_mix/internal/model"

	"gorm.io/gorm"
)

type BlogService struct {
	DB *gorm.DB
}

func (s *BlogService) CreateArticle(article *model.Article) error {
	return s.DB.Create(article).Error
}
