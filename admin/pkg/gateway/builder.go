package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/pkg/starter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strings"
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

	// swagger 文档
	swagger        bool
	swaggerPrefix  string
	swaggerHandler http.Handler
}

func NewGWBuild(options ...Options) *GWBuild {
	gw := &GWBuild{}
	for _, opt := range options {
		opt.apply(gw)
	}

	gw.mux = runtime.NewServeMux()
	gw.RouterGroup = RouterGroup{mux: gw.mux, gw: gw}
	gw.comp = starter.NewComp("Gateway", gw.conf.RunMode == DevMod)
	gw.comp.SetCompItem("port", gw.conf.Address)

	// 拨号连接信息设置
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if gw.conf.GrpcTransData > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(gw.conf.GrpcTransData),
			grpc.MaxCallRecvMsgSize(gw.conf.GrpcTransData),
		))
	}

	// 连接grpc
	conn, err := grpc.DialContext(context.Background(), gw.conf.GrpcAddress, opts...)
	if err != nil {
		panic("拨号失败:" + err.Error())
	}
	gw.conn = conn

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
	// 启动时进行注册
	if build.swagger {
		build.comp.SetCompItem(build.swaggerPrefix, build.swaggerPrefix)
	}

	// 服务注册
	build.httpServer = &http.Server{
		Addr: build.conf.Address,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 跨域配置
			method := r.Method
			origin := r.Header.Get("Origin")
			if origin != "" {
				// 允许来源配置
				w.Header().Set("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
				w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
				w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			if method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// 启动swagger配置信息
			if build.swagger && strings.HasPrefix(r.URL.Path, build.swaggerPrefix) {
				build.swaggerHandler.ServeHTTP(w, r)
				return
			}

			build.mux.ServeHTTP(w, r)
		}),
	}

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
