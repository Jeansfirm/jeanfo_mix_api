package main

import "fmt"

type BaseReq struct{}

type ReqAutoFiller interface {
	AutoFill()
}

type PageReq struct {
	BaseReq
	Page     int `uri:"Page" json:"Page" form:"Page" binding:"omitempty,min=1"`
	PageSize int `uri:"PageSize" json:"PageSize" form:"PageSize" binding:"omitempty,min=1"`
}

func (pr *PageReq) AutoFill() {
	if pr.Page == 0 {
		pr.Page = 1
	}
	if pr.PageSize == 0 {
		pr.PageSize = 15
	}
}

type ListArticleReq struct {
	PageReq

	UserID int `json:"UserID"`
}

// AutoFill: SpecificReq 覆盖 BaseReqData 的 AutoFill 方法
// func (s *ListArticleReq) AutoFill() {
// 	s.PageReq.AutoFill()
// 	fmt.Println("SpecificReq AutoFill called")
// }

func main() {
	// 初始化结构体
	base := &PageReq{}
	req := &ListArticleReq{}

	b, ok := any(base).(*BaseReq)
	fmt.Println(b, ok)

	r, ok := any(req).(*PageReq)
	fmt.Println(r, ok)

	r2, ok := any(req).(*ListArticleReq)
	fmt.Println(r2, ok)

	r3, ok := any(req).(ReqAutoFiller)
	fmt.Println(r3, ok)
	(r3).AutoFill()
}
