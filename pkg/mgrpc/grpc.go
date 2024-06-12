package mgrpc

import (
	"fmt"
	"github.com/zxcblog/rat-race/framework/logger"
	"google.golang.org/grpc"
	"net"
)

type GrpcConf struct {
	Port string
}

type Grpc struct {
	server *grpc.Server
	listen net.Listener

	log logger.ILogger
}

func New(conf GrpcConf, log logger.ILogger) *Grpc {
	server := grpc.NewServer()
	lis, err := net.Listen("tcp", conf.Port)
	if err != nil {
		panic(fmt.Sprintf("grpc启动监听端口失败：%s", err.Error()))
	}
	log.InfoF("grpc启动监听端口：%s", conf.Port)

	return &Grpc{
		server: server,
		listen: lis,
		log:    log,
	}
}

func (g *Grpc) Run() {
	serverInfo := g.server.GetServiceInfo()
	for pkg, info := range serverInfo {
		g.log.DebugF("%-6s %-25s", pkg, info.Methods)
	}

	go func() {
		if err := g.server.Serve(g.listen); err != nil {
			panic(fmt.Sprintf("grpc服务启动失败%s", err.Error()))
		}
	}()
}

// RegisterServer 注册grpc服务
func (g *Grpc) RegisterServer(desc *grpc.ServiceDesc, imp any) {
	g.server.RegisterService(desc, imp)
}

func (g *Grpc) Close() {
	g.server.GracefulStop()
}
