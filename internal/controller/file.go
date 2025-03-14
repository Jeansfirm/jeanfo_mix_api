package controller

import (
	"jeanfo_mix/internal/definition"
	"jeanfo_mix/internal/service"
	context_util "jeanfo_mix/util/context"
	request_util "jeanfo_mix/util/request"
	response_util "jeanfo_mix/util/response"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	Service *service.FileService
}

func (c *FileController) UploadFile(gtx *gin.Context) {
	httpContext := context_util.NewHttpContext(gtx)
	userID := httpContext.ClientData().UserID

	file, err := gtx.FormFile("file")
	if err != nil {
		response_util.NewResponse(gtx).SetMsg("fail to get file from form: " + err.Error()).FailBadRequest()
		return
	}
	fileName := file.Filename

	src, err := file.Open()
	if err != nil {
		response_util.NewResponse(gtx).SetMsg("fail to open file from form: " + err.Error()).FailBadRequest()
		return
	}
	defer src.Close()

	metaID, relativePath, err := c.Service.UploadFile(src, fileName, uint(userID), true)
	if err != nil {
		response_util.NewResponse(gtx).SetMsg("fail to save file: " + err.Error()).FailBadRequest()
		return
	}

	response_util.NewResponse(gtx).SetData(
		map[string]any{"MetaID": metaID, "RelativePath": relativePath},
	).Success()
}

// @Summary Common: Download File
// @Tags Common
// @Param query query definition.DownloadFileReq true "download file"
// @Router /api/file/download [get]
// @Security BearerAuth
func (c *FileController) DownloadFile(gtx *gin.Context) {
	userID := context_util.NewHttpContext(gtx).SessionData().UserID
	req := request_util.NewRequest[definition.DownloadFileReq](gtx)

	filePath, fileName, err := c.Service.DownloadFile(uint(userID), req.Data.MetaID)
	if err != nil {
		response_util.NewResponse(gtx).SetMsg(err.Error()).FailBadRequest()
		return
	}

	gtx.FileAttachment(filePath, fileName)
}
