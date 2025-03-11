package service

import (
	"errors"
	"fmt"
	user_definition "jeanfo_mix/internal/definition/user"
	"jeanfo_mix/internal/model"

	"gorm.io/gorm"
)

func (us *UserService) Get(userId int) (*model.User, error) {
	user := &model.User{}
	err := us.DB.First(user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户未找到")
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (us *UserService) List() {

}

func (us *UserService) Update(req *user_definition.UpdateUserReq) error {
	user := &model.User{}
	err := us.DB.First(user, req.UserID).Error
	if err != nil {
		return errors.New("user not found")
	}

	if req.AvatarRelativePath != "" {
		user.AvatarRelativePath = req.AvatarRelativePath
	}

	err = us.DB.Save(user).Error
	if err != nil {
		return fmt.Errorf("update user fail: %s", err.Error())
	}

	return nil
}
