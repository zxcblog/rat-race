package router

import (
	"github.com/zxcblog/rat-race/internal/client"
	"github.com/zxcblog/rat-race/internal/router/pb/user"
	user2 "github.com/zxcblog/rat-race/internal/server/user"
	"github.com/zxcblog/rat-race/pkg/grpc"
	grpc2 "google.golang.org/grpc"
)

func GRPCRouter() *grpc.GRPCBuild {
	server := grpc.NewGRPCBuild(grpc.WithConfig(client.GrpcConf), grpc.WithLogOptions(client.Log))

	// 注册grpc服务
	server.RegisterServer(func(s *grpc2.Server) {
		user.RegisterUserServer(s, user2.NewUserServer())
	})

	return server
}
