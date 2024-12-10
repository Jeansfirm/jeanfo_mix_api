package service

// import (
// 	"errors"
// 	"time"

// 	"context"
// 	"jeanfo_mix/config"
// 	"jeanfo_mix/internal/model"
// 	"jeanfo_mix/util"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/google/uuid"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// // 配置
// var (
// 	jwtSecret   = []byte(config.AppConfig.JWTSecret) // 从配置文件读取
// 	redisClient = redis.NewClient(&redis.Options{
// 		Addr:     config.AppConfig.Redis.Addr,
// 		Password: config.AppConfig.Redis.Password,
// 		DB:       config.AppConfig.Redis.DB,
// 	})
// 	ctx = context.Background()
// )

// // HashPassword 加密密码
// func HashPassword(password string) (string, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

// // VerifyPassword 验证密码
// func VerifyPassword(hashedPassword, password string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
// }

// // GenerateToken 生成JWT Token
// func GenerateToken(userID string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"user_id": userID,
// 		"exp":     time.Now().Add(72 * time.Hour).Unix(), // 3天有效期
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtSecret)
// }

// // SaveSession 保存用户会话信息到Redis
// func SaveSession(userID string, token string) error {
// 	return redisClient.Set(ctx, "session:"+userID, token, 72*time.Hour).Err()
// }

// // Register 注册用户
// func Register(db *gorm.DB, username, password string) (*model.User, error) {
// 	if !util.IsValidPassword(password) {
// 		return nil, errors.New("密码必须包含大写、小写字母和数字")
// 	}

// 	var existingUser model.User
// 	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
// 		return nil, errors.New("用户名已存在")
// 	}

// 	hashedPassword, err := HashPassword(password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user := &model.User{
// 		ID:           uuid.New().String(),
// 		Username:     username,
// 		PasswordHash: hashedPassword,
// 	}
// 	if err := db.Create(user).Error; err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// // Login 登录用户
// func Login(db *gorm.DB, username, password string) (*model.User, string, error) {
// 	var user model.User
// 	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
// 		return nil, "", errors.New("用户名或密码错误")
// 	}

// 	if !VerifyPassword(user.PasswordHash, password) {
// 		return nil, "", errors.New("用户名或密码错误")
// 	}

// 	token, err := GenerateToken(user.ID)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	// 保存到Redis
// 	if err := SaveSession(user.ID, token); err != nil {
// 		return nil, "", err
// 	}

// 	return &user, token, nil
// }

// // ThirdPartyLogin 第三方登录
// func ThirdPartyLogin(db *gorm.DB, provider, providerID string) (*model.User, string, error) {
// 	var user model.User
// 	if err := db.Where("provider = ? AND provider_id = ?", provider, providerID).First(&user).Error; err != nil {
// 		// 如果用户不存在，创建新用户
// 		user = model.User{
// 			ID:         uuid.New().String(),
// 			Provider:   provider,
// 			ProviderID: providerID,
// 		}
// 		if err := db.Create(&user).Error; err != nil {
// 			return nil, "", err
// 		}
// 	}

// 	token, err := GenerateToken(user.ID)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	// 保存到Redis
// 	if err := SaveSession(user.ID, token); err != nil {
// 		return nil, "", err
// 	}

// 	return &user, token, nil
// }
