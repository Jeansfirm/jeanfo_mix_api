package controller

import (
	user_service "jeanfo_mix/internal/service/user"
	"jeanfo_mix/util"

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

// auth

func (uc *UserController) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "params error: " + err.Error()
		util.NewResponse(ctx).SetMsg(msg).FailBadRequest()
		return
	}

	user, err := uc.Service.Register(req.RType, req.UserName, req.Password, "", "", "", "")
	if err != nil {
		msg := "register fail: " + err.Error()
		util.NewResponse(ctx).SetMsg(msg).FailBadRequest()
		return
	}

	util.NewResponse(ctx).SetMsg("register success").SetData(user).Success()
}

func (uc *UserController) Login(ctx *gin.Context) {

}

func (uc *UserController) Logout(ctx *gin.Context) {

}

func (uc *UserController) ChangePasswd(ctx *gin.Context) {

}

// user crud

func (uc *UserController) Get(ctx *gin.Context) {

}

func (uc *UserController) List(ctx *gin.Context) {

}
