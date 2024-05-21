package render

import "net/http"

type Render interface {
	// Render 自定义返回信息
	Render(w http.ResponseWriter) error
	// WriteContentType 写入自定义的ContentType
	WriteContentType(w http.ResponseWriter)
}

var (
	_ Render = (*JSON)(nil)
)

func writerContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = val
	}
}
