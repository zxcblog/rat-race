package gateway

import (
	"github.com/zxcblog/rat-race/framework/gateway/render"
	"net/http"
	"time"
)

// Context gateway 自定义路由上下文解析
type Context struct {
	Writer  responseWriter
	Request *http.Request

	PathParams map[string]string

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request, pathParams map[string]string) *Context {
	ctx := &Context{
		Writer:     responseWriter{w, 0},
		Request:    r,
		PathParams: pathParams,
		Path:       r.URL.Path,
		Method:     r.Method,
	}

	return ctx
}

// JSON 返回json格式数据信息
func (c *Context) JSON(code int, obj any) {
	c.Render(code, render.JSON{Data: obj})
}

func (c *Context) Render(code int, r render.Render) {
	c.Status(code)

	if err := r.Render(c.Writer); err != nil {
		panic(err.Error())
	}
}

// Status 设置 http 返回的状态值
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

// File 启动一个文件服务
func (c *Context) File(filepath string) {
	http.ServeFile(c.Writer, c.Request, filepath)
}

// FileFromFS 通过 http.FileSystem 启动一个文件流服务
// filepath: 路径参数名， 从请求路径中获取到filepath的地址信息，将信息当作路由
func (c *Context) FileFromFS(filepath string, fs http.FileSystem) {
	defer func(old string) {
		c.Request.URL.Path = old
	}(c.Request.URL.Path)

	c.Request.URL.Path = c.PathParams[filepath]

	http.FileServer(fs).ServeHTTP(c.Writer, c.Request)
}

// 实现 Context.Context 接口

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Request.Context().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.Request.Context().Done()
}

func (c *Context) Err() error {
	return c.Request.Context().Err()
}

func (c *Context) Value(key any) any {
	return c.Request.Context().Value(key)
}
