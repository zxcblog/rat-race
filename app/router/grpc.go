package router

import (
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/zxcblog/rat-race/app/client"
	"github.com/zxcblog/rat-race/app/pb/user"
	"github.com/zxcblog/rat-race/app/router/middleware"
	user2 "github.com/zxcblog/rat-race/app/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strings"
)

func GrpcRouter() {
	server := grpc.NewServer(
		// 一元RPC 拦截器注册
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(middleware.RecoveryInterceptor()), // panic 保护
		),
	)

	// 监听端口
	lis, err := net.Listen("tcp", client.Config.Server.GrpcConf.Port)
	if err != nil {
		panic(fmt.Sprintf("grpc启动监听端口失败：%s", err.Error()))
	}
	client.Logger.DebugF("grpc启动监听端口：%s", client.Config.Server.GrpcConf.Port)

	// 注册服务
	user.RegisterUserServer(server, user2.NewUserServer())

	// 打印注册的服务信息，并注册到etcd中
	reflection.Register(server)
	serverInfo := server.GetServiceInfo()
	pkgs := make([]string, 0, len(serverInfo))
	for pkg := range serverInfo {
		pkgs = append(pkgs, pkg)
	}
	client.Logger.DebugF("%-6s :%-25s", "成功注册服务", strings.Join(pkgs, ","))

	grpcproxy.Re

	// 启动grpc服务
	go func() {
		if err = server.Serve(lis); err != nil {
			panic(fmt.Sprintf("grpc服务启动失败：%s", err.Error()))
		}
	}()

	// 运行并注册关闭服务
	client.Shutdown.Register(func() {
		server.GracefulStop()
	})
}
