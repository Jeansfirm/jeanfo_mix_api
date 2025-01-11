package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ginCtx *gin.Context

	HttpCode int

	Code int
	Msg  string
	Data interface{}
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		ginCtx: ctx, HttpCode: http.StatusOK,
	}
}

func (r *Response) Send() *Response {
	// todo data判断转换成json

	r.ginCtx.JSON(r.HttpCode,
		gin.H{
			"Code": r.Code,
			"Msg":  r.Msg,
			"Data": r.Data,
		},
	)
	return r
}

func (r *Response) SetHttpCode(httpCode int) *Response {
	r.HttpCode = httpCode
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) SetMsg(msg string) *Response {
	r.Msg = msg
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) Success() *Response {
	return r.SetHttpCode(http.StatusOK).SetCode(0).Send()
}

func (r *Response) Fail() *Response {
	return r.SetCode(-1).Send()
}

func (r *Response) FailBadRequest() *Response {
	return r.SetHttpCode(http.StatusBadRequest).Fail()
}

func (r *Response) FailInternalServerError() *Response {
	return r.SetHttpCode(http.StatusInternalServerError).Fail()
}
