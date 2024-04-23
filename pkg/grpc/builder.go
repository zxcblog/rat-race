package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// GRPCBuild 服务启动
type GRPCBuild struct {
	grpcS  *grpc.Server
	config *Config

	errs []error
}

// NewGRPCBuild 初始化grpc服务
func NewGRPCBuild(config *Config) *GRPCBuild {
	builder := &GRPCBuild{config: config}

	// 初始化grpc服务
	// TODO 用户自定义注册拦截器
	builder.grpcS = grpc.NewServer()

	// dev环境添加反射
	if builder.config.RunMode == DevMod {
		reflection.Register(builder.grpcS)
	}

	return builder
}

func (build *GRPCBuild) RegisterServer(opts ...func(s *grpc.Server)) *GRPCBuild {
	for _, opt := range opts {
		opt(build.grpcS)
	}
	return build
}

// Start 服务启动
func (build *GRPCBuild) Start() error {
	// 启动grpc服务
	lis, err := net.Listen("tcp", build.config.Address)
	if err != nil {
		return err
	}

	if err := build.grpcS.Serve(lis); err != nil {
		return err
	}
	return nil
}
