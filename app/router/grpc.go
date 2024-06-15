package router

import (
	"github.com/zxcblog/rat-race/app/client"
	"github.com/zxcblog/rat-race/app/pb/user"
	user2 "github.com/zxcblog/rat-race/app/service/user"
	"github.com/zxcblog/rat-race/pkg/mgrpc"
	"google.golang.org/grpc"
)

func GrpcRouter() {
	grpcInstance := mgrpc.New(client.Config.Server.GrpcConf, client.Logger, client.Etcd)

	// 注册运行服务
	grpcInstance.RegisterServer(func(s grpc.ServiceRegistrar) {
		user.RegisterUserServer(s, user2.NewUserServer())
	})

	// 运行并注册关闭服务
	grpcInstance.Run()
	client.Shutdown.Register(func() {
		grpcInstance.Close()
	})
}
