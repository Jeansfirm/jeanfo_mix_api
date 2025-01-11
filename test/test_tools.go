package test

import (
	"encoding/json"
	"io"
	reponse_util "jeanfo_mix/util/response"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var GTestTool TestTool
var initOnce sync.Once

type TestTool struct {
	db     *gorm.DB
	router *gin.Engine
}

func init() {
	initOnce.Do(func() {
		db, err := SetupTestDBV2()
		if err != nil {
			panic("setup test db fail: " + err.Error())
		}
		GTestTool.db = db
		GTestTool.router = SetupTestRouter(db)
	})
}

type HttpCall struct {
	T    *testing.T
	Req  *http.Request
	Resp *httptest.ResponseRecorder
}

func (tt *TestTool) GenHttpCall(t *testing.T, method, url string, body io.Reader) *HttpCall {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	return &HttpCall{T: t, Req: req, Resp: resp}
}

func (tt *HttpCall) Run() {
	GTestTool.router.ServeHTTP(tt.Resp, tt.Req)
}

type RespData = reponse_util.ResponsePayload

func (tt *HttpCall) GetRespData() *RespData {
	var respData RespData
	err := json.Unmarshal(tt.Resp.Body.Bytes(), &respData)
	assert.Nil(tt.T, err, "parse resp body to json fail")

	return &respData
}
