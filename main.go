package main

import (
	"fmt"
	"github.com/zxcblog/rat-race/internal/router/pb/user"
	user2 "github.com/zxcblog/rat-race/internal/server/user"
	"github.com/zxcblog/rat-race/pkg/grpc"
	grpc2 "google.golang.org/grpc"
)

func main() {
	err := grpc.NewGRPCBuild().RegisterServer(func(s *grpc2.Server) {
		user.RegisterUserServer(s, user2.NewUserServer())
	}).Start()
	fmt.Println(err)
}
