package request_util

import (
	error_definition "jeanfo_mix/internal/definition/error"

	"github.com/gin-gonic/gin"
)

type Request[T any] struct {
	ginCtx *gin.Context
	Data   *T
}

func NewRequest[T any](ctx *gin.Context) *Request[T] {
	var req Request[T]
	req.ginCtx = ctx

	var reqData T
	req.Data = &reqData

	paramsErr := error_definition.BadRequestError{}
	paramsErr.Code = -2
	paramsErr.Msg = "params error: "

	// err = ctx.ShouldBindUri(&reqData)
	// if err == nil {
	// 	log_util.Debug("params from uri parse")
	// 	return &req
	// } else if err != io.EOF {
	// 	paramsErr.Msg += err.Error() + " - parsing from uri"
	// 	panic(err)
	// }

	// err = ctx.ShouldBindJSON(&reqData)
	// if err == nil {
	// 	log_util.Debug("params from json parse")
	// 	return &req
	// } else if err != io.EOF {
	// 	paramsErr.Msg += err.Error() + " - parsing from json"
	// 	panic(err)
	// }

	err := ctx.ShouldBind(&reqData)
	if err != nil {
		paramsErr.Msg += err.Error()
		panic(err)
	}

	return &req
}
