package middleware

import (
	"strings"

	auth_service "jeanfo_mix/internal/service/auth"
	response_util "jeanfo_mix/util/response"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 验证用户登录态
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader, authParam := "", ""
		tokenPart := ""

		authHeader = c.GetHeader("Authorization")

		if authHeader != "" {
			authParts := strings.Split(authHeader, " ")
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				response_util.NewResponse(c).SetMsg("auth header not bear type").FailUnauthorized()
				c.Abort()
				return
			}
			tokenPart = authParts[1]
		} else {
			authParam = c.Query("_Auth_Token")
			tokenPart = authParam
		}

		if tokenPart == "" {
			response_util.NewResponse(c).SetMsg("no auth header or param found").FailUnauthorized()
			c.Abort()
			return
		}

		token := auth_service.ClientToken(tokenPart)
		clientData := &auth_service.ClientData{}
		if err := clientData.Load(token); err != nil {
			response_util.NewResponse(c).SetMsg("token parse fail: " + err.Error()).FailUnauthorized()
			c.Abort()
			return
		}

		sessData, err := clientData.GetSessionData()
		if err != nil {
			response_util.NewResponse(c).SetMsg("get session from token fail: " + err.Error()).FailUnauthorized()
			c.Abort()
			return
		}

		c.Set("ClientData", clientData)
		c.Set("SessionData", sessData)

		c.Next()
	}
}
