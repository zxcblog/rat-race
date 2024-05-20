package main

import (
	"context"
	"github.com/zxcblog/rat-race/internal/client"
	"github.com/zxcblog/rat-race/internal/router"
	"github.com/zxcblog/rat-race/pkg/starter"
	"github.com/zxcblog/rat-race/pkg/tools"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
)

// TODO
// 跨域配置
// gateway返回数据格式统一设置
// 请求时记录访问日志
// swagger配置信息设置
func main() {
	if err := client.Init("./config/config.yaml"); err != nil {
		log.Fatalf("初始化全局信息失败：%s", err.Error())
		return
	}

	grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)

	//grpclog.SetLoggerV2(client.Log)

	// 启动grpc
	grpcServer := router.GRPCRouter()
	grpcServer.Start()

	// 启动gateway
	gwServer, err := router.GWRouter()
	if err != nil {
		client.Log.ErrorF(context.Background(), "gateway启动失败：%s", err.Error())
		return
	}
	gwServer.Start()

	// 输出配置信息
	starter.Print()

	// 服务关闭
	tools.NewShutDown().Close(
		// 关闭 http 服务
		func() {
			gwServer.ShutDown()
		},

		// 关闭 grpc 服务
		func() {
			grpcServer.ShutDown()
		},

		func() {
			client.Close()
		},
	)
}
