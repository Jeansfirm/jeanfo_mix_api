package controller

import (
	"jeanfo_mix/internal/model"
	auth_service "jeanfo_mix/internal/service/auth"
	user_service "jeanfo_mix/internal/service/user"
	context_util "jeanfo_mix/util/context"
	reponse_util "jeanfo_mix/util/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *user_service.UserService
}

type RegisterReq struct {
	RType    user_service.RegisterType `json:"RType" bind:"required"`
	UserName string                    `json:"UserName"`
	Password string                    `json:"Password"`
}

type LoginReq struct {
	LType    user_service.LoginType `json:"LType" bind:"required"`
	UserName string                 `json:"UserName"`
	Password string                 `json:"Password"`
}

type LoginResp struct {
	Token string
	User  model.User
}

// auth

func (uc *UserController) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "params error: " + err.Error()
		reponse_util.NewResponse(ctx).SetMsg(msg).FailBadRequest()
		return
	}

	user, err := uc.Service.Register(req.RType, req.UserName, req.Password, "", "", "", "")
	if err != nil {
		msg := "register fail: " + err.Error()
		reponse_util.NewResponse(ctx).SetMsg(msg).FailBadRequest()
		return
	}

	reponse_util.NewResponse(ctx).SetMsg("register success").SetData(user).Success()
}

func (uc *UserController) Login(ctx *gin.Context) {
	var req LoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		reponse_util.NewResponse(ctx).SetMsg("params error: " + err.Error()).FailBadRequest()
		return
	}
	user, clientToken, err := uc.Service.Login(req.LType, req.UserName, req.Password)
	if err != nil {
		reponse_util.NewResponse(ctx).SetMsg("login fail: " + err.Error()).FailBadRequest()
		return
	}

	resp := LoginResp{Token: clientToken, User: *user}
	reponse_util.NewResponse(ctx).SetMsg("login success").SetData(resp).Success()
}

func (uc *UserController) Logout(ctx *gin.Context) {
	httpContext := context_util.NewHttpContext(ctx)
	clientData := httpContext.ClientData()
	clientToken, _ := clientData.GetToken()
	auth_service.LogoutUser(clientToken)

	reponse_util.NewResponse(ctx).SetMsg("logout success").Success()
}

func (uc *UserController) ChangePasswd(ctx *gin.Context) {

}

// user crud

func (uc *UserController) Get(ctx *gin.Context) {

}

func (uc *UserController) List(ctx *gin.Context) {

}
