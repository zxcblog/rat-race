package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// GRPCBuild 服务启动
type GRPCBuild struct {
	ctx      context.Context
	grpcS    *grpc.Server
	svrName  string
	svrAddr  string
	etcdAddr string

	errs []error
}

// NewGRPCBuild 初始化grpc服务
func NewGRPCBuild() *GRPCBuild {
	builder := &GRPCBuild{
		ctx: context.Background(),
	}

	// 初始化grpc服务
	// TODO 用户自定义注册拦截器
	builder.grpcS = grpc.NewServer()

	// TODO 添加判断是否需要添加反射信息
	reflection.Register(builder.grpcS)

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
	lis, err := net.Listen("tcp", ":6666")
	if err != nil {
		return err
	}

	if err := build.grpcS.Serve(lis); err != nil {
		return err
	}
	return nil
}
