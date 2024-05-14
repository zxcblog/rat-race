package grpc

import (
	"github.com/zxcblog/rat-race/pkg/logger"
	"github.com/zxcblog/rat-race/pkg/starter"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

// GRPCBuild 服务启动
type GRPCBuild struct {
	grpcS  *grpc.Server
	config *Config

	listen net.Listener
	errs   []error

	interceptors []grpc.UnaryServerInterceptor
	log          logger.ILogger
	opts         []grpc.ServerOption

	compItem []starter.CompItem
}

// NewGRPCBuild 初始化grpc服务
func NewGRPCBuild(options ...OptionFunc) *GRPCBuild {
	builder := &GRPCBuild{
		interceptors: make([]grpc.UnaryServerInterceptor, 0),
		opts:         make([]grpc.ServerOption, 0),
	}
	for _, opt := range options {
		opt.apply(builder)
	}

	// 设置发送和请求接收数据大小
	if builder.config.TransDataSize > 0 {
		builder.opts = append(builder.opts, grpc.MaxSendMsgSize(builder.config.TransDataSize), grpc.MaxRecvMsgSize(builder.config.TransDataSize))
	}

	if len(builder.interceptors) > 0 {
		builder.opts = append(builder.opts, grpc.ChainUnaryInterceptor(builder.interceptors...))
	}

	// 初始化grpc服务
	builder.grpcS = grpc.NewServer(builder.opts...)

	// 启动grpc服务
	lis, err := net.Listen("tcp", builder.config.Address)
	if err != nil {
		panic("端口监听失败")

	}

	builder.listen = lis
	return builder
}

func (build *GRPCBuild) RegisterServer(opts ...func(s *grpc.Server)) *GRPCBuild {
	for _, opt := range opts {
		opt(build.grpcS)
	}
	return build
}

// Start 服务启动
func (build *GRPCBuild) Start() {
	serverInfo := build.grpcS.GetServiceInfo()

	build.SetCompItem("port", build.config.Address)
	for pkg, info := range serverInfo {

		names := make([]string, len(info.Methods))
		for i, method := range info.Methods {
			names[i] = method.Name
		}
		build.SetCompItem(pkg, strings.Join(names, "\n"))
	}

	starter.RegisterComp(build)
	go func() {
		if err := build.grpcS.Serve(build.listen); err != nil {
			// 打印日志
			log.Fatalf("服务启动失败", err.Error())
		}
	}()
}

func (build *GRPCBuild) ShutDown() {
	build.grpcS.GracefulStop()
}

// CompName 实现IComp来输出配置信息
func (build *GRPCBuild) CompName() string {
	return "GRPC"
}

func (build *GRPCBuild) GetCompItem() []starter.CompItem {
	return build.compItem
}

func (build *GRPCBuild) SetCompItem(key, val string) {
	if len(build.compItem) < 0 {
		build.compItem = make([]starter.CompItem, 0, 10)
	}
	build.compItem = append(build.compItem, starter.CompItem{Key: key, Value: val})
}

func (build *GRPCBuild) IsDev() bool {
	return build.config.RunMode == DevMod
}
