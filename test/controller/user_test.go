package controller_test

import (
	"bytes"
	"jeanfo_mix/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserController_Retrieve(t *testing.T) {
	testUserName := "kkkk"

	httpCall := test.GTestTool.GenHttpCall(t, "GET", "/api/user", bytes.NewBuffer([]byte{}))
	httpCall.LoginAs(testUserName)
	httpCall.Run()
	assert.Equal(t, http.StatusOK, httpCall.Resp.Code)
	respData := httpCall.GetRespData()
	userInfo := respData.Data.(map[string]any)
	assert.Equal(t, testUserName, userInfo["Username"].(string))
}
