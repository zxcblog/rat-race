package mgateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/framework/logger"
	"google.golang.org/protobuf/encoding/protojson"
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
		//Origim

		//// 信息认证，因此需要允许部分header穿越gw
		//runtime.WithIncomingHeaderMatcher(gwHeaderHandler),
		//

		//runtime.WithErrorHandler(ErrHandler), // 统一错误处理

		// 此处只能用来追加Header信息， 不能覆盖返回结构体, 可以通过将返回抛到错误信息中进行返回
		//runtime.WithForwardResponseOption(ResponseHandler),

		// 处理响应JSON问题
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true, // 按照PB协议定义命名
				EmitUnpopulated: true, // null的数据返回，不过滤
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
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
