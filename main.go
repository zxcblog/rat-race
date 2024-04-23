package main

import (
	"fmt"
	"github.com/zxcblog/rat-race/config"
	"github.com/zxcblog/rat-race/internal/router/pb/user"
	user2 "github.com/zxcblog/rat-race/internal/server/user"
	"github.com/zxcblog/rat-race/pkg/grpc"
	grpc2 "google.golang.org/grpc"
	"log"
)

func main() {
	// 初始化配置信息
	if err := config.InitConfig("./config/config.yaml"); err != nil {
		log.Fatalf("初始化配置项失败：%s", err.Error())
		return
	}

	err := grpc.NewGRPCBuild(config.GrpcConf).RegisterServer(func(s *grpc2.Server) {
		user.RegisterUserServer(s, user2.NewUserServer())
	}).Start()
	fmt.Println(err)
}
