package service

import (
	"jeanfo_mix/internal/model"

	"gorm.io/gorm"
)

type DemoService struct {
	DB *gorm.DB
}

func (s *DemoService) GetHelloWord() string {
	return "hello world"
}

func (s *DemoService) GetDemos(title string, page, pageSize int) ([]model.Demo, error) {
	var demos []model.Demo
	query := s.DB
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&demos).Error

	return demos, err
}
