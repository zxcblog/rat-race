package main

import (
	"github.com/zxcblog/rat-race/config"
	"github.com/zxcblog/rat-race/internal/client"
	"github.com/zxcblog/rat-race/internal/router"
	"github.com/zxcblog/rat-race/pkg/logger"
	"github.com/zxcblog/rat-race/pkg/tools"
	"log"
)

func main() {
	// 初始化配置信息
	if err := config.InitConfig("./config/config.yaml"); err != nil {
		log.Fatalf("初始化配置项失败：%s", err.Error())
		return
	}

	// 启动日志服务
	client.Log = logger.NewLogger(config.LogConf)

	// 启动grpc
	grpcServer := router.GRPCRouter()
	grpcServer.Start()

	// 启动gateway
	gwServer := router.GWRouter()
	gwServer.Start()

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

		// 关闭数据库连接
		func() {
			client.DB.Close()
		},

		// 关闭日志
		func() {
			if err := client.Log.Close(); err != nil {
				log.Fatalf("日志服务关闭失败：%s", err.Error())
			}
		},
	)
}
