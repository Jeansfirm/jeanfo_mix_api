package controller_test

import (
	"bytes"
	"encoding/json"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/test"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemoController_CreateDemo(t *testing.T) {
	db := test.SetupTestDB(t)
	router := test.SetupTestRouter(db)

	payload := map[string]string{"title": "Test Title", "content": "Test Content"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/demos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, "Test Title", response["data"].(map[string]interface{})["Title"])

	var demo model.Demo
	db.First(&demo)
	assert.Equal(t, "Test Content", demo.Content)
}

func TestDemoController_DeleteDemo(t *testing.T) {
	db := test.SetupTestDB(t)
	router := test.SetupTestRouter(db)

	demo := &model.Demo{Title: "Test", Content: "Content"}
	db.Create(demo)

	req, _ := http.NewRequest("DELETE", "/api/demos/"+strconv.Itoa(int(demo.ID)), nil)
	resp := httptest.NewRecorder()

	// fmt.Printf("resp: %v, req: %v", resp, req)
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var storedDemo model.Demo
	result := db.First(&storedDemo, demo.ID)
	assert.Error(t, result.Error)
	assert.Equal(t, "record not found", result.Error.Error())
}
