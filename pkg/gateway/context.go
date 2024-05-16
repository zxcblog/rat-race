package gateway

import (
	"github.com/zxcblog/rat-race/pkg/gateway/render"
	"net/http"
)

type H map[string]interface{}

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
	c.Render(render.JSON{Data: obj})
}

func (c *Context) Render(r render.Render) {
	if err := r.Render(c.Writer); err != nil {
		panic(err.Error())
	}
}
