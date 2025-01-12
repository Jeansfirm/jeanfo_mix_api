package context_util

import (
	"fmt"
	auth_service "jeanfo_mix/internal/service/auth"
	session_util "jeanfo_mix/util/session"

	"github.com/gin-gonic/gin"
)

type HttpContext struct {
	GinCtx *gin.Context
}

func NewHttpContext(ginCtx *gin.Context) *HttpContext {
	return &HttpContext{GinCtx: ginCtx}
}

func (hc *HttpContext) SessionData() *session_util.SessionData {
	value, ok := hc.GinCtx.Get("SessionData")
	if !ok {
		return nil
	}
	sessData, ok := value.(*session_util.SessionData)
	if !ok {
		fmt.Println("not correct SessionData type from ginContext")
		return nil
	}
	return sessData
}

func (hc *HttpContext) ClientData() *auth_service.ClientData {
	value, ok := hc.GinCtx.Get("ClientData")
	if !ok {
		return nil
	}
	clientData, ok := value.(*auth_service.ClientData)
	if !ok {
		fmt.Println("not correct ClientData type from ginContext")
		return nil
	}
	return clientData
}
