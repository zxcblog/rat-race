package mgateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/framework/logger"
	"github.com/zxcblog/rat-race/pkg/mgateway/handler"
	"net/http"
)

type GatewayConf struct {
	Port string
}

type HandlerFunc func(ctx *Context)

type H map[string]interface{}

type Gateway struct {
	RouterGroup

	log  logger.ILogger
	conf GatewayConf

	mux    *runtime.ServeMux
	server *http.Server
}

// New 启动gateway服务
func New(config GatewayConf, log logger.ILogger) *Gateway {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(handler.ErrHandler),
	)
	//https://github.com/janrs-io/Jgrpc-response

	gateway := &Gateway{
		RouterGroup: RouterGroup{
			basePath: "/",
			mux:      mux,
			log:      log,
		},
		mux:  mux,
		log:  log,
		conf: config,
	}

	return gateway
}

func (g *Gateway) Run() {
	g.server = &http.Server{
		Addr:    g.conf.Port,
		Handler: g.mux,
	}
	g.log.InfoF("gateway服务启动成功，监听地址信息：%s", g.server.Addr)

	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			panic(fmt.Sprintf("gateway 服务启动失败:%s", err.Error()))
		}
	}()
}

func (g *Gateway) Close(ctx context.Context) error {
	return g.server.Shutdown(ctx)
}

func (g *Gateway) GetServerMux() *runtime.ServeMux {
	return g.mux
}
