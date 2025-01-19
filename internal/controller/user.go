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
	RType    string `json:"RType" bind:"required"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type LoginReq struct {
	LType    string `json:"LType" bind:"required"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type LoginResp struct {
	Token string
	User  model.User
}

// auth

// @Summary Auth: Register
// @Tags Auth
// @Param register body RegisterReq true "register"
// @Router /api/auth/register [post]
func (uc *UserController) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "params error: " + err.Error()
		reponse_util.NewResponse(ctx).SetMsg(msg).FailBadRequest()
		return
	}

	user, err := uc.Service.Register(user_service.RegisterType(req.RType), req.UserName, req.Password, "", "", "", "")
	if err != nil {
		msg := "register fail: " + err.Error()
		reponse_util.NewResponse(ctx).SetMsg(msg).FailBadRequest()
		return
	}

	reponse_util.NewResponse(ctx).SetMsg("register success").SetData(user).Success()
}

// @Summary Auth: Login
// @Tags Auth
// @Param login body LoginReq true "login"
// @Router /api/auth/login [post]
func (uc *UserController) Login(ctx *gin.Context) {
	var req LoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		reponse_util.NewResponse(ctx).SetMsg("params error: " + err.Error()).FailBadRequest()
		return
	}
	user, clientToken, err := uc.Service.Login(user_service.LoginType(req.LType), req.UserName, req.Password)
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
	sessData := context_util.NewHttpContext(ctx).SessionData()
	user := uc.Service.GetUser(sessData.UserName)

	reponse_util.NewResponse(ctx).SetData(user).Success()
}

func (uc *UserController) List(ctx *gin.Context) {

}
