package grpc

import (
	"github.com/zxcblog/rat-race/pkg/logger"
	"google.golang.org/grpc"
)

type Options interface {
	apply(build *GRPCBuild)
}

type OptionFunc func(*GRPCBuild)

func (f OptionFunc) apply(b *GRPCBuild) {
	f(b)
}

func WithConfig(config *Config) OptionFunc {
	return func(build *GRPCBuild) {
		build.config = config
	}
}

func WithInterceptors(interceptors ...grpc.UnaryServerInterceptor) OptionFunc {
	return func(build *GRPCBuild) {
		build.interceptors = append(build.interceptors, interceptors...)
	}
}

func WithServerOptions(options ...grpc.ServerOption) OptionFunc {
	return func(build *GRPCBuild) {
		build.opts = append(build.opts)
	}
}

func WithLogOptions(logger logger.ILogger) OptionFunc {
	return func(build *GRPCBuild) {
		build.log = logger
	}
}
