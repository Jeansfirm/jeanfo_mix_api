package middleware

import (
	error_definition "jeanfo_mix/internal/definition/error"
	"jeanfo_mix/util/log_util"
	reponse_util "jeanfo_mix/util/response"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoverMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resp := reponse_util.NewResponse(c)

				expected := false
				msg := "An Unknown Error Occured"

				switch err := err.(type) {
				case error_definition.BadRequestError:
					resp.SetCode(err.Code).SetMsg(err.Error()).FailBadRequest()
					expected = true
				case error:
					msg = err.Error()
				case string:
					msg = err
				}

				if !expected {
					resp.SetCode(-99).SetMsg(msg).FailInternalServerError()
					stack := debug.Stack()
					log_util.Error("RecoverMiddleWare recover panic from:%v\nStack Trace:\n%s\n", err, stack)
				}

				c.Abort()
			}
		}()

		c.Next()
	}
}
