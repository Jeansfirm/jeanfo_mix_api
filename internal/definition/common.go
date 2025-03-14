package definition

type DownloadFileReq struct {
	BaseReq

	MetaID string `json:"MetaID" binding:"required"`
}
