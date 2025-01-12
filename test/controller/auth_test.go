package controller_test

import (
	"bytes"
	"encoding/json"
	"jeanfo_mix/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	correctUserName = "jeanfo"
	correctPassword = "123AbcEf45x"
)

func TestUserController_Main(t *testing.T) {
	// test register
	payload := map[string]string{
		"RType": "Normal", "UserName": correctUserName, "Password": correctPassword,
	}
	body, _ := json.Marshal(payload)

	rHttpCall := test.GTestTool.GenHttpCall(t, "POST", "/api/auth/register", bytes.NewBuffer(body))
	rHttpCall.Run()
	assert.Equal(t, http.StatusOK, rHttpCall.Resp.Code)
	rRespData := rHttpCall.GetRespData()
	assert.Equal(t, rRespData.Code, 0)

	// test login
	payload = map[string]string{"LType": "Normal", "UserName": correctUserName, "Password": correctPassword}
	lBody, _ := json.Marshal(payload)
	lHttpCall := test.GTestTool.GenHttpCall(t, "POST", "/api/auth/login", bytes.NewBuffer(lBody))
	lHttpCall.Run()
	assert.Equal(t, http.StatusOK, lHttpCall.Resp.Code)
	lRespData := lHttpCall.GetRespData()
	assert.NotEmpty(t, lRespData.Data.(map[string]any)["Token"])

	// test logout
}

func TestUserController_FailAction(t *testing.T) {
	// register fail
	rPayLoad := map[string]string{"RType": "Normal", "UserName": "kk", "Password": "123"}
	rBody, _ := json.Marshal(rPayLoad)
	rHttpCall := test.GTestTool.GenHttpCall(t, "POST", "/api/auth/register", bytes.NewBuffer(rBody))
	rHttpCall.Run()
	assert.Equal(t, rHttpCall.Resp.Code, http.StatusBadRequest)
	rRespData := rHttpCall.GetRespData()
	assert.Contains(t, rRespData.Msg, "用户名长度必须在")

	// login fail
	lPayload := map[string]string{"LType": "Normal", "UserName": "jeanfo", "Password": "xxx"}
	bodyBytes, _ := json.Marshal(lPayload)
	httpCall := test.GTestTool.GenHttpCall(t, "POST", "/api/auth/login", bytes.NewBuffer(bodyBytes))
	httpCall.Run()
	assert.Equal(t, httpCall.Resp.Code, http.StatusBadRequest)
	respData := httpCall.GetRespData()
	assert.Equal(t, respData.Code, -1)

	// logout fail
	lfHttpCall := test.GTestTool.GenHttpCall(t, "POST", "/api/auth/logout", bytes.NewBuffer([]byte{}))
	lfHttpCall.Run()
	assert.Equal(t, lfHttpCall.Resp.Code, http.StatusUnauthorized)
	lfRespData := lfHttpCall.GetRespData()
	assert.Contains(t, lfRespData.Msg, "no auth header found")

	lfHttpCall2 := test.GTestTool.GenHttpCall(t, "POST", "/api/auth/logout", bytes.NewBuffer([]byte{}))
	lfHttpCall2.Req.Header.Set("Ahuthorization", "xx ee")
	lfHttpCall2.Run()
	assert.Equal(t, lfHttpCall2.Resp.Code, http.StatusUnauthorized)
	lfRespData2 := lfHttpCall2.GetRespData()
	assert.Contains(t, lfRespData2.Msg, "not bear type")
}
