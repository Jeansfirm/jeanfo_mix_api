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
		authString := c.GetHeader("Ahuthorization")
		if authString == "" {
			response_util.NewResponse(c).SetMsg("no auth header found").FailUnauthorized()
			c.Abort()
			return
		}

		authParts := strings.Split(authString, " ")
		if len(authParts) != 2 || authParts[0] != "Bear" {
			response_util.NewResponse(c).SetMsg("auth header not bear type").FailUnauthorized()
			c.Abort()
			return
		}

		token := auth_service.ClientToken(authParts[1])
		clientData := &auth_service.ClientData{}
		if err := clientData.Load(token); err != nil {
			response_util.NewResponse(c).SetMsg("auth header parse fail: " + err.Error()).FailUnauthorized()
			c.Abort()
			return
		}

		sessData, err := clientData.GetSessionData()
		if err != nil {
			response_util.NewResponse(c).SetMsg("get session from auth header fail: " + err.Error()).FailUnauthorized()
			c.Abort()
			return
		}

		c.Set("ClientData", clientData)
		c.Set("SessionData", sessData)

		c.Next()
	}
}
