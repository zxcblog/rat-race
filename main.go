package main

import (
	"github.com/zxcblog/rat-race/internal/client"
	"github.com/zxcblog/rat-race/internal/router"
	"github.com/zxcblog/rat-race/pkg/starter"
	"github.com/zxcblog/rat-race/pkg/tools"
	"log"
)

func main() {
	if err := client.Init("./config/config.yaml"); err != nil {
		log.Fatalf("初始化全局信息失败：%s", err.Error())
		return
	}

	// 启动grpc
	grpcServer := router.GRPCRouter()
	grpcServer.Start()

	// 启动gateway
	gwServer := router.GWRouter()
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
