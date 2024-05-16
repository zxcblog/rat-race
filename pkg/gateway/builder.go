package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/pkg/starter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type GWBuild struct {
	RouterGroup

	// mux grpc-gateway 的请求多路复用器，
	// 所有信息根据该方法进行实现
	mux        *runtime.ServeMux
	httpServer *http.Server

	conn *grpc.ClientConn
	conf *Config

	comp starter.IComp
}

func NewGWBuild(conf *Config) *GWBuild {
	gw := &GWBuild{
		conf: conf,
	}

	gw.mux = runtime.NewServeMux()
	gw.RouterGroup = RouterGroup{mux: gw.mux, gw: gw}
	gw.comp = starter.NewComp("Gateway", conf.RunMode == DevMod)
	gw.comp.SetCompItem("port", conf.Address)

	// 拨号连接信息设置
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if conf.GrpcTransData > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(conf.GrpcTransData),
			grpc.MaxCallRecvMsgSize(conf.GrpcTransData),
		))
	}

	// 连接grpc
	conn, err := grpc.DialContext(context.Background(), conf.GrpcAddress, opts...)
	if err != nil {
		panic("拨号失败:" + err.Error())
	}
	gw.conn = conn

	// 服务注册
	gw.httpServer = &http.Server{Addr: conf.Address, Handler: gw.mux}
	return gw
}

func (build *GWBuild) RegisterServer(funcs ...func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) *GWBuild {
	for _, f := range funcs {
		if err := f(context.Background(), build.mux, build.conn); err != nil {
			log.Fatalf("服务注册失败:%s", err.Error())
		}
	}

	return build
}

func (build *GWBuild) Start() {
	go func() {
		if err := build.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("gateway 服务启动失败:%s", err.Error())
		}
	}()
}

func (build *GWBuild) ShutDown() {
	ctx, cancel := context.WithTimeout(context.Background(), build.conf.ShutDownTime)
	defer cancel()

	if err := build.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("服务关闭失败")
	}
}
