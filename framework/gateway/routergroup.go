package gateway

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/framework/logger"
	"net/http"
)

// IRoutes 定义所有路由器句柄接口
type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	GET(string, HandlerFunc) IRoutes
	POST(string, HandlerFunc) IRoutes
	DELETE(string, HandlerFunc) IRoutes
	PATCH(string, HandlerFunc) IRoutes
	PUT(string, HandlerFunc) IRoutes
	OPTIONS(string, HandlerFunc) IRoutes
	HEAD(string, HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	//Static(string, string) IRoutes
	//StaticFS(string, http.FileSystem) IRoutes
}

// RouterGroup 内部路由器配置
type RouterGroup struct {
	// 用来进行路由分组
	basePath string

	// 路由中间件
	Handlers []HandlerFunc

	mux *runtime.ServeMux
	log logger.ILogger
}

func (r *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	r.Handlers = append(r.Handlers, middleware...)
	return r
}

// Group 路由分组
func (r *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		basePath: relativePath,
		Handlers: handlers,
		mux:      r.mux,
	}
}

// 合并路由中间件
func (r *RouterGroup) combineHandlers(handlers ...HandlerFunc) []HandlerFunc {
	finalSize := len(r.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {
		panic("中间件太多了")
	}

	// 使用copy进行切片复制会比使用for循环快很多，而且避免了内存分配，所以性能会更好
	mergedHandlers := make([]HandlerFunc, finalSize)
	copy(mergedHandlers, r.Handlers)
	copy(mergedHandlers[len(r.Handlers):], handlers)
	return mergedHandlers
}

func (r *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(r.basePath, relativePath)
}

func (r *RouterGroup) addHandle(httpMethod, relativePath string, handlers HandlerFunc) IRoutes {
	r.log.DebugF("%-6s %-25s --> %s (%d handlers)", httpMethod, relativePath, nameOfFunction(handlers), len(r.Handlers)+1)

	// 将分组路径进行合并
	relativePath = joinPaths(r.basePath, relativePath)
	err := r.mux.HandlePath(httpMethod, relativePath, r.encapHandleFunc(handlers))
	if err != nil {
		panic(err.Error())
	}
	return r
}

// encapHandleFunc 因为是将请求路由注册到mux中，所以不用实现路由树
// 为了放便注册路由信息和对请求的参数信息和返回信息进行处理,
// 对 runtime.HandlerFunc 进行二次封装处理
func (r *RouterGroup) encapHandleFunc(handlerFunc HandlerFunc) runtime.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx := newContext(w, req, pathParams, append(r.Handlers, handlerFunc))
		ctx.Next()
	}
}

func (r *RouterGroup) GET(path string, handler HandlerFunc) IRoutes {
	return r.addHandle(http.MethodGet, path, handler)
}
func (r *RouterGroup) POST(path string, handler HandlerFunc) IRoutes {
	return r.addHandle(http.MethodPost, path, handler)
}
func (r *RouterGroup) DELETE(path string, handler HandlerFunc) IRoutes {
	return r.addHandle(http.MethodDelete, path, handler)
}
func (r *RouterGroup) PATCH(path string, handler HandlerFunc) IRoutes {
	return r.addHandle(http.MethodPatch, path, handler)
}
func (r *RouterGroup) PUT(path string, handler HandlerFunc) IRoutes {
	return r.addHandle(http.MethodPut, path, handler)
}
func (r *RouterGroup) OPTIONS(path string, handler HandlerFunc) IRoutes {
	return r.addHandle(http.MethodOptions, path, handler)
}
func (r *RouterGroup) HEAD(path string, handler HandlerFunc) IRoutes {
	return r.addHandle(http.MethodHead, path, handler)
}

func (r *RouterGroup) StaticFile(relativePath, filepath string) IRoutes {
	return r.GET(relativePath, func(ctx *Context) {
		ctx.File(filepath)
	})
}

func (r *RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes {
	return r.GET(relativePath, func(ctx *Context) {
		ctx.FileFromFS(filepath, fs)
	})
}
