package definition

// Base def

type BaseReq struct{}

type ReqAutoFiller interface {
	AutoFill()
}

func (*BaseReq) AutoFill() {}

// PageReq def

type PageReq struct {
	BaseReq
	Page     int `uri:"Page" json:"Page" form:"Page" binding:"omitempty,min=1"`
	PageSize int `uri:"PageSize" json:"PageSize" form:"PageSize" binding:"omitempty,min=1"`
}

func (pr *PageReq) AutoFill() {
	pr.BaseReq.AutoFill()

	if pr.Page == 0 {
		pr.Page = 1
	}
	if pr.PageSize == 0 {
		pr.PageSize = 15
	}
}
