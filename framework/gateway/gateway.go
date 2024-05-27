package gateway

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/framework/logger"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type H map[string]interface{}

type Gateway struct {
	RouterGroup

	log logger.ILogger

	mux    *runtime.ServeMux
	server *http.Server
}

// New 启动gateway服务
func New(log logger.ILogger) *Gateway {
	mux := runtime.NewServeMux()

	gateway := &Gateway{
		RouterGroup: RouterGroup{
			basePath: "/",
			mux:      mux,
			log:      log,
		},
		mux: mux,
		log: log,
	}

	return gateway
}

func (g *Gateway) Run() {
	g.server = &http.Server{
		Addr:    ":8080",
		Handler: g.mux,
	}
	g.log.InfoF("gateway服务启动成功，监听地址信息：%s", g.server.Addr)

	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			panic(fmt.Sprintf("gateway 服务启动失败:%s", err.Error()))
		}
	}()
}
