package user_definition

import "jeanfo_mix/internal/definition"

type UpdateUserBaseReq struct {
	definition.BaseReq
	AvatarRelativePath string `json:"AvatarRelativePath"`
}

type UpdateUserReq struct {
	UpdateUserBaseReq
	UserID uint `json:"UserID" binding:"required"`
}

type UpdateUserMyReq struct {
	UpdateUserBaseReq
	UserID uint `json:"UserID"`
}
