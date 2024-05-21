package gateway

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type H map[string]interface{}

type Gateway struct {
	RouterGroup

	mux    *runtime.ServeMux
	server *http.Server
}

// New 启动gateway服务
func New() *Gateway {
	mux := runtime.NewServeMux()

	return &Gateway{
		RouterGroup: RouterGroup{
			basePath: "/",
			mux:      mux,
		},
		mux: mux,
	}
}

func (g *Gateway) Run() {
	g.server = &http.Server{
		Addr:    ":8080",
		Handler: g.mux,
	}

	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			log.Fatalf("gateway 服务启动失败:%s", err.Error())
		}
	}()
}
