package service

import (
	"errors"
	"jeanfo_mix/internal/model"
	auth_service "jeanfo_mix/internal/service/auth"
	"jeanfo_mix/util"

	"gorm.io/gorm"
)

type RegisterType string
type LoginType string

const (
	RegisterTypeNormal     RegisterType = "Normal"
	RegisterTypePhone      RegisterType = "Phone"
	RegisterTypeThirdParty RegisterType = "ThirdParty"

	LoginTypeNormal     LoginType = LoginType(RegisterTypeNormal)
	LoginTypePhone      LoginType = LoginType(RegisterTypePhone)
	LoginTypeThirdParth LoginType = LoginType(RegisterTypeThirdParty)
)

const DefaultPassword string = "--Empty--"

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) Register(rType RegisterType,
	username, password, phone, provider, providerID, providerToken string,
) (*model.User, error) {
	switch rType {
	case RegisterTypeNormal:
		if len(username) == 0 && len(password) == 0 {
			return nil, errors.New("必须同时指定用户名和密码")
		}
		return us.RegisterNormal(username, password)
	case RegisterTypePhone:
		// todo 手机号必须
		return us.RegisterPhone(phone)
	case RegisterTypeThirdParty:
		// todo  第三方信息必须
		return us.RegisterThirdParty(provider, providerID, providerToken)
	default:
	}
	return nil, errors.New("非法注册类型: " + string(rType))
}

func (us *UserService) RegisterNormal(username string, password string) (*model.User, error) {
	user := &model.User{
		Username:     username,
		RegisterType: string(RegisterTypeNormal),
	}

	return us.CreateUser(user, password, true)
}

func (us *UserService) RegisterPhone(phone string) (*model.User, error) {
	user := &model.User{
		Username:     auth_service.GenerateUserName(),
		RegisterType: string(RegisterTypePhone),
	}

	return us.CreateUser(user, DefaultPassword, false)
}

func (us *UserService) RegisterThirdParty(provider, providerID, providerToken string) (*model.User, error) {
	user := &model.User{
		Username:      auth_service.GenerateUserName(),
		RegisterType:  string(RegisterTypeThirdParty),
		Provider:      provider,
		ProviderID:    providerID,
		ProviderToken: providerToken,
	}

	return us.CreateUser(user, DefaultPassword, false)
}

// CreateUser 注册用户
func (us *UserService) CreateUser(user *model.User, password string, saveHashedPassword bool) (*model.User, error) {
	// 验证用户名格式
	if len(user.Username) < 3 || len(user.Username) > 20 {
		return nil, errors.New("用户名长度必须在3到20个字符之间")
	}

	// 验证密码强度
	if password != DefaultPassword && !util.IsValidPassword(password) {
		return nil, errors.New("密码必须包含大写、小写字母和数字，长度8-20")
	}

	// 检查用户名是否已存在
	var existingUser model.User
	if err := us.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 如果指定手机号，检查手机号是否存在
	// todo

	// 如果指定第三方登录，检查第三方信息是否存在
	// todo

	// 加密密码
	hashedPassword := password
	if saveHashedPassword {
		_hpwd, err := auth_service.HashPassword(password)
		if err != nil {
			return nil, err
		}
		hashedPassword = _hpwd
	}

	// 创建用户
	user.PasswordHash = hashedPassword
	if err := us.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Login 登录用户
func (us *UserService) Login(lType LoginType,
	username, password string,
) (*model.User, auth_service.ClientToken, error) {
	var user model.User
	if err := us.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, "", errors.New("用户名或密码错误")
	}

	if !auth_service.VerifyPassword(user.PasswordHash, password) {
		return nil, "", errors.New("用户名或密码错误")
	}

	sessionData := auth_service.SessionData{UserID: user.ID, UserName: user.Username, Role: user.Role}
	jwt, err := auth_service.LoginUser(&sessionData)

	return &user, jwt, err
}

func (us *UserService) Logout(jwt auth_service.ClientToken) error {
	err := auth_service.LogoutUser(jwt)
	return err
}

// ChangePassword 修改密码
func (us *UserService) ChangePassword(userID uint, oldPassword, newPassword string, verifyOldPassword bool) error {
	var user model.User
	if err := us.DB.First(&user, userID).Error; err != nil {
		return errors.New("用户不存在")
	}

	// 验证新密码强度
	if !util.IsValidPassword(newPassword) {
		return errors.New("新密码必须包含大写、小写字母和数字，长度8-20")
	}

	// 如果是第三方注册用户且不需要验证旧密码
	if user.PasswordHash == DefaultPassword {
		hashedPassword, err := auth_service.HashPassword(newPassword)
		if err != nil {
			return err
		}
		return us.DB.Model(&user).Update("password_hash", hashedPassword).Error
	}

	// 验证旧密码
	if !auth_service.VerifyPassword(user.PasswordHash, oldPassword) {
		return errors.New("旧密码不正确")
	}

	// 更新密码
	hashedPassword, err := auth_service.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return us.DB.Model(&user).Update("password_hash", hashedPassword).Error
}
