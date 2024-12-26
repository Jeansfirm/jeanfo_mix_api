package util

import (
	"errors"
	"time"

	"jeanfo_mix/config"

	"github.com/golang-jwt/jwt/v4"
)

// 使用给定data生成token字符串
//
//	validSeconds 有效期，单位 秒
func JwtGenerateToken(data map[string]interface{}, validSeconds ...int) (string, error) {
	claims := jwt.MapClaims(data)
	if _, ok := claims["exp"]; !ok {
		seconds := 3 * 24 * 60 * 60
		if len(validSeconds) > 0 {
			seconds = validSeconds[0]
		}
		claims["exp"] = time.Now().Add(time.Duration(seconds) * time.Second).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Printf("jwt token object: %v \njwt secret: %v \n\n", token, config.AppConfig.JWTSecret)

	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// ParseToken 验证并解析JWT
//
// Note: 如果token已过期，返回错误“无效的token”
func JwtParseToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("无效的Token")
	}
	return claims, nil
}

// TokenExpired 验证Token是否过期
func JwtTokenExpired(token string) bool {
	claims, err := JwtParseToken(token)
	if err != nil {
		return true
	}

	if exp, ok := claims["exp"].(float64); ok {
		return time.Now().Unix() > int64(exp)
	}

	return true
}
