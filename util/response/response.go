package reponse_util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ginCtx   *gin.Context
	HttpCode int

	ResponsePayload
}

type ResponsePayload struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data any    `json:"Data"`
}

type PaginatedData struct {
	Page     int   `json:"Page"`
	PageSize int   `json:"PageSize"`
	Total    int64 `json:"Total"`
	Rows     any   `json:"Rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	resp := &Response{
		ginCtx: ctx, HttpCode: http.StatusOK,
	}

	return resp
}

var New = NewResponse

func (r *Response) Send() *Response {
	if r.ginCtx.Writer.Written() {
		// todo 需要清空原有的Body，否则会导致Body为两个json串的拼接，这里尚未实现
		r.ginCtx.Writer.WriteHeader(r.HttpCode)
		body, _ := json.Marshal(r.ResponsePayload)
		r.ginCtx.Writer.Write(body) // 会导致叠串
	} else {
		r.ginCtx.JSON(r.HttpCode, r.ResponsePayload)
	}
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

func (r *Response) SetDataPaginated(total int64, rows any, page, pageSize int) *Response {
	r.Data = PaginatedData{
		Total: total, Rows: rows, Page: page, PageSize: pageSize,
	}
	return r
}

func (r *Response) Success() *Response {
	return r.SetHttpCode(http.StatusOK).SetCode(0).Send()
}

func (r *Response) Fail() *Response {
	if r.Code == 0 {
		r.Code = -1
	}
	return r.Send()
}

func (r *Response) FailUnauthorized() *Response {
	return r.SetHttpCode(http.StatusUnauthorized).Fail()
}

func (r *Response) FailBadRequest() *Response {
	return r.SetHttpCode(http.StatusBadRequest).Fail()
}

func (r *Response) FailInternalServerError() *Response {
	return r.SetHttpCode(http.StatusInternalServerError).Fail()
}
