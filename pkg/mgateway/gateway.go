package mgateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/framework/logger"
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
	mux := runtime.NewServeMux()

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

//// WithGrpcCfgConnOpt 增加构建选项，读取config中gw_port的端口配置请求grpc
//// 和WithGrpcConnOpt的区别是不用手动拨号，直接获取conn完成连接，
//// 其他拨号形式请使用ConnOpt, 并搭配上有效的Conn
//func WithGrpcCfgConnOpt(ctx context.Context, optFunc func(mux *runtime.ServeMux, cfgConn *grpc.ClientConn) error) GWBuildOption {
//	return func(builder *GWBuilder) {
//		conn, connErr := dial(ctx, "tcp", web_info.GetSvrAddr())
//		if connErr != nil {
//			panic(connErr)
//		}
//		if err := optFunc(builder.mux, conn); err != nil {
//			log.GetLogger().ErrorF(context.Background(), "gRPC-Gateway注册Grpc连接发现错误[WithCfgConn], error: %s", err.Error())
//		}
//	}
//}
