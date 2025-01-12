package middleware

import (
	"fmt"
	reponse_util "jeanfo_mix/util/response"
	"runtime/debug"

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

				stack := debug.Stack()
				fmt.Printf("RecoverMiddleWare recover panic from:%v\nStack Trace:\n%s\n", err, stack)
			}
		}()

		c.Next()
	}
}
