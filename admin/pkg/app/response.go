package app

//
//import (
//	"github.com/zxcblog/rat-race/pkg/gateway"
//	"net/http"
//)
//
//type Response struct {
//	ctx *gateway.Context
//}
//
//type Pager struct {
//	Page      int `json:"page"`
//	PageSize  int `json:"page_size"`
//	TotalRows int `json:"total_rows"`
//}
//
//func NewResponse(ctx *gateway.Context) *Response {
//	return &Response{ctx: ctx}
//}
//
//func (r *Response) ToResponse(data interface{}) {
//	if data == nil {
//		data = gateway.H{}
//	}
//
//	r.ctx.JSON(http.StatusOK, data)
//}
//
//func (r *Response) ToResponseList(list interface{}, page, pageSize, totalRows int) {
//	r.ctx.JSON(http.StatusOK, gateway.H{
//		"list": list,
//		"page": gateway.H{
//			"page":       page,
//			"page_size":  pageSize,
//			"total_rows": totalRows,
//		},
//	})
//}
//
//func (r *Response) ToResponseError(err *Error) {
//	res := gateway.H{"code": err.Code(), "msg": err.Msg()}
//	details := err.Details()
//	if len(details) > 0 {
//		res["details"] = details
//	}
//	r.ctx.JSON(err.StatusCode(), res)
//}
