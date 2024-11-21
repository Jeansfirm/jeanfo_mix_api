package service

import (
	"errors"
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

func (s *DemoService) CreateDemo(title string, content string) (*model.Demo, error) {
	demo := &model.Demo{
		Title:   title,
		Content: content,
	}

	if err := s.DB.Create(demo).Error; err != nil {
		return nil, err
	}

	return demo, nil
}

func (s *DemoService) DeleteDemo(id uint) error {
	result := s.DB.Delete(&model.Demo{}, id)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return result.Error
}
