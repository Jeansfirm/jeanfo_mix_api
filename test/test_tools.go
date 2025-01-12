package test

import (
	"encoding/json"
	"fmt"
	"io"
	auth_service "jeanfo_mix/internal/service/auth"
	user_service "jeanfo_mix/internal/service/user"
	reponse_util "jeanfo_mix/util/response"
	session_util "jeanfo_mix/util/session"
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
	db     *gorm.DB
	router *gin.Engine

	T    *testing.T
	Req  *http.Request
	Resp *httptest.ResponseRecorder
}

func (tt *TestTool) GenHttpCall(t *testing.T, method, url string, body io.Reader) *HttpCall {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	return &HttpCall{db: tt.db, router: tt.router, T: t, Req: req, Resp: resp}
}

func (tt *HttpCall) Run() {
	GTestTool.router.ServeHTTP(tt.Resp, tt.Req)
}

func (tt *HttpCall) Reset() {
	tt.Resp = httptest.NewRecorder()
}

type RespData = reponse_util.ResponsePayload

func (tt *HttpCall) GetRespData() *RespData {
	var respData RespData
	err := json.Unmarshal(tt.Resp.Body.Bytes(), &respData)
	assert.Nil(tt.T, err, "parse resp body to json fail")

	return &respData
}

func (hc *HttpCall) LoginAs(userName string) {
	userService := user_service.UserService{DB: hc.db}
	user := userService.GetUser(userName)
	if user == nil {
		newUser, err := userService.RegisterNormal(userName, "123Abc@Ef56")
		if err != nil {
			panic(fmt.Sprintf("Register %s fail: %s", userName, err.Error()))
		}
		user = newUser
	}

	sessionData := session_util.SessionData{UserID: user.ID, UserName: userName}
	clientToken, _ := auth_service.LoginUser(&sessionData)

	hc.Req.Header.Set("Authorization", "Bearer "+clientToken)
}
