package main

import (
	"github.com/zxcblog/rat-race/admin/pkg/tools"
	"github.com/zxcblog/rat-race/framework"
	"github.com/zxcblog/rat-race/framework/gateway"
	"github.com/zxcblog/rat-race/framework/logger"
)

// TODO
// 跨域配置
// gateway返回数据格式统一设置
// 请求时记录访问日志
// swagger配置信息设置
func main() {
	engine := framework.New(
		framework.WithLogInstance(logger.NewLogger("rat-race", logger.WithConsoleCore(logger.GetZapEncode(), "debug"))),
	)

	engine.GET("/hello", func(ctx *gateway.Context) {
		ctx.JSON(404, "hello word")
	})

	//engine.StaticFile("/config", "./config/config.yaml")
	engine.StaticFileFS("/config/{path}", "path", gateway.Dir("./config", true))

	engine.Run()

	//if err := client.Init("./config/config.yaml"); err != nil {
	//	log.Fatalf("初始化全局信息失败：%s", err.Error())
	//	return
	//}
	//
	//grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	//
	////grpclog.SetLoggerV2(client.Log)
	//
	//// 启动grpc
	//grpcServer := router.GRPCRouter()
	//grpcServer.Start()
	//
	//// 启动gateway
	//gwServer, err := router.GWRouter()
	//if err != nil {
	//	client.Log.ErrorF(context.Background(), "gateway启动失败：%s", err.Error())
	//	return
	//}
	//gwServer.Start()
	//
	//// 输出配置信息
	//starter.Print()
	//
	// 服务关闭
	tools.NewShutDown().Close()
}
