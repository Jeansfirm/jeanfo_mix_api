package middleware

import (
	reponse_util "jeanfo_mix/util/response"

	"github.com/gin-gonic/gin"
)

func RecoverMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := "An Unknown Error Occured"

				switch err := err.(type) {
				case error:
					msg = err.Error()
				case string:
					msg = err
				}

				reponse_util.NewResponse(c).SetMsg(msg).FailInternalServerError()
				c.Abort()
			}
		}()

		c.Next()
	}
}
