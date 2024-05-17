package router

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/internal/client"
	"github.com/zxcblog/rat-race/internal/router/pb/user"
	"github.com/zxcblog/rat-race/internal/server/captcha"
	"github.com/zxcblog/rat-race/pkg/gateway"
	grpc2 "google.golang.org/grpc"
)

func GWRouter() *gateway.GWBuild {
	server := gateway.NewGWBuild(client.GwConf).RegisterServer(
		func(ctx context.Context, mux *runtime.ServeMux, conn *grpc2.ClientConn) error {
			if err := user.RegisterUserHandler(ctx, mux, conn); err != nil {
				return err
			}
			return nil
		},
	)

	// 验证码操作
	catpchaServer := captcha.NewCaptchaServer()
	server.GET("/v1/captcha", catpchaServer.Get)
	return server
}
