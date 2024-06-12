package router

import (
	"github.com/zxcblog/rat-race/app/client"
	"github.com/zxcblog/rat-race/pkg/mgrpc"
)

func GrpcRouter() {
	grpcInstance := mgrpc.New(client.Config.Server.GrpcConf, client.Logger)

	grpcInstance.Run()
}
