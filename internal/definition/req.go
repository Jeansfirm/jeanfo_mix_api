package definition

import "gorm.io/gorm"

// Base def

type BaseReq struct{}

type BaseUserReq struct {
	BaseReq
	UserID uint `json:"UserID"`
}

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

type PageUserReq struct {
	PageReq
	UserID uint `json:"UserID"`
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

func (pr *PageReq) Paginate(query *gorm.DB) (int64, *gorm.DB) {
	var total int64
	query.Count(&total)

	offset := (pr.Page - 1) * pr.PageSize
	query = query.Offset(offset).Limit(pr.PageSize)

	return total, query
}
