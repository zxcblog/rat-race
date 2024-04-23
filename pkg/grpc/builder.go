package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// GRPCBuild 服务启动
type GRPCBuild struct {
	grpcS  *grpc.Server
	config *Config

	listen net.Listener
	errs   []error
}

// NewGRPCBuild 初始化grpc服务
func NewGRPCBuild(config *Config) *GRPCBuild {
	builder := &GRPCBuild{config: config}

	opts := make([]grpc.ServerOption, 0)

	// 设置发送和请求接收数据大小
	if config.TransDataSize > 0 {
		opts = append(opts, grpc.MaxSendMsgSize(int(config.TransDataSize)), grpc.MaxRecvMsgSize(int(config.TransDataSize)))
	}

	// 初始化grpc服务
	// TODO 用户自定义注册拦截器
	builder.grpcS = grpc.NewServer(opts...)

	// dev环境添加反射
	if builder.config.RunMode == DevMod {
		reflection.Register(builder.grpcS)
	}

	// 启动grpc服务
	lis, err := net.Listen("tcp", config.Address)
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
