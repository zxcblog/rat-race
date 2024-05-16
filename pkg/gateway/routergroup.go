package gateway

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

// IRoutes 定义所有路由器句柄接口
type IRoutes interface {
	GET(string, HandlerFunc) IRoutes
	POST(string, HandlerFunc) IRoutes
}

// RouterGroup 内部路由器配置
type RouterGroup struct {
	// mux grpc-gateway 的请求多路复用器，结合mux实现自定义接口信息
	mux *runtime.ServeMux
	gw  *GWBuild
}

func (r *RouterGroup) addHandle(httpMethod, relativePath string, handlers HandlerFunc) {
	// 注册自定义添加的路由信息
	r.gw.comp.SetCompItem(httpMethod, relativePath)

	r.mux.HandlePath(httpMethod, relativePath, r.encapHandleFunc(handlers))
}

// encapHandleFunc 因为是将请求路由注册到mux中，所以不用实现路由树
// 为了放便注册路由信息和对请求的参数信息和返回信息进行处理,
// 对 runtime.HandlerFunc 进行二次封装处理
func (r *RouterGroup) encapHandleFunc(handlerFunc HandlerFunc) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		handlerFunc(newContext(w, r, pathParams))
	}
}

func (r *RouterGroup) GET(s string, handlerFunc HandlerFunc) IRoutes {
	r.addHandle(http.MethodGet, s, handlerFunc)
	return r
}

func (r *RouterGroup) POST(s string, handlerFunc HandlerFunc) IRoutes {
	r.addHandle(http.MethodPost, s, handlerFunc)
	return r
}
