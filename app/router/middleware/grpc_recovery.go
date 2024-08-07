package middleware

import (
	"context"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/zxcblog/rat-race/app/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
		client.Logger.ErrorFWithCtx(ctx, "发生panic错误：%s", p)
		return status.Errorf(codes.Internal, "panic:%v", p)
	})
}
