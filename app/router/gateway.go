package router

import "C"
import (
	"context"
	"github.com/zxcblog/rat-race/app/client"
	"github.com/zxcblog/rat-race/app/service/captcha"
	"github.com/zxcblog/rat-race/pkg/mgateway"
	"time"
)

func GatewayRouter() {
	route := mgateway.New(client.Config.Server.GatewayConf, client.Logger)

	// 注册自定义请求地址
	captS := captcha.NewCaptchaServer()
	route.GET("/v1/captch/img", captS.Get)

	// 注册grpc服务
	//route.

	route.Run()
	// 注册关闭服务
	client.Shutdown.Register(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := route.Close(ctx); err != nil {
			client.Logger.ErrorF("gateway服务关闭失败:", err)
		}

		client.Logger.InfoF("gateway服务退出")
	})
}
