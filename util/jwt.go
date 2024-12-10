package util

// import (
// 	"errors"
// 	"time"

// 	"jeanfo_mix/config"

// 	"github.com/golang-jwt/jwt/v4"
// )

// // ParseToken 验证并解析JWT
// func ParseToken(tokenString string) (jwt.MapClaims, error) {
// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(config.Config.JWTSecret), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, errors.New("无效的Token")
// 	}
// 	return claims, nil
// }

// // TokenExpired 验证Token是否过期
// func TokenExpired(token string) bool {
// 	claims, err := ParseToken(token)
// 	if err != nil {
// 		return true
// 	}

// 	if exp, ok := claims["exp"].(float64); ok {
// 		return time.Now().Unix() > int64(exp)
// 	}

// 	return true
// }
