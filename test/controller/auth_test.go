package controller_test

import (
	"bytes"
	"encoding/json"
	"jeanfo_mix/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	correctUserName = "jeanfo"
	correctPassword = "123AbcEf45x"
)

func TestUserController_Main(t *testing.T) {
	db := test.SetupTestDB(t)
	router := test.SetupTestRouter(db)

	// test register
	payload := map[string]string{
		"RType": "Normal", "UserName": correctUserName, "Password": correctPassword,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	var respData map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &respData)
	assert.Nil(t, err)
	assert.Equal(t, respData["Code"], float64(0))

	// test login
	payload = map[string]string{"LType": "Normal", "UserName": correctUserName, "Password": correctPassword}
	body, _ = json.Marshal(payload)
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	respData = map[string]any{}
	err = json.Unmarshal(resp.Body.Bytes(), &respData)
	assert.Nil(t, err)
	data := respData["Data"]
	assert.NotEmpty(t, data.(map[string]any)["Token"])
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
}
