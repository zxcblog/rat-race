package mgrpc

import (
	"fmt"
	"github.com/zxcblog/rat-race/framework/logger"
	"github.com/zxcblog/rat-race/pkg/metcd"
	"google.golang.org/grpc"
	"net"
	"strings"
)

type GrpcConf struct {
	Port       string
	Host       string
	ServerName string
}

func (g GrpcConf) Addr() string {
	return g.Host + ":" + g.Port
}

type Grpc struct {
	server *grpc.Server
	listen net.Listener

	config GrpcConf
	log    logger.ILogger
	metcd  *metcd.MEtcd
}

func New(conf GrpcConf, log logger.ILogger, etcd *metcd.MEtcd) *Grpc {
	server := grpc.NewServer()
	lis, err := net.Listen("tcp", conf.Addr())
	if err != nil {
		panic(fmt.Sprintf("grpc启动监听端口失败：%s", err.Error()))
	}
	log.InfoF("grpc启动监听端口：%s", conf.Addr())

	return &Grpc{
		server: server,
		listen: lis,
		config: conf,
		log:    log,
		metcd:  etcd,
	}
}

func (g *Grpc) Run() {
	serverInfo := g.server.GetServiceInfo()
	pkgs := make([]string, 0, len(serverInfo))
	for pkg := range serverInfo {
		pkgs = append(pkgs, pkg)
	}
	g.log.DebugF("%-6s :%-25s", "成功注册服务", strings.Join(pkgs, ","))

	go func() {
		if err := g.server.Serve(g.listen); err != nil {
			panic(fmt.Sprintf("grpc服务启动失败%s", err.Error()))
		}
	}()

	// 获取当前ip，端口号，将实现的服务注册到etcd中
	if err := g.register(pkgs); err != nil {
		panic(fmt.Sprintf("grpc服务注册etcd失败：%s", err.Error()))
	}
}

// RegisterServer 注册grpc服务
func (g *Grpc) RegisterServer(f func(s grpc.ServiceRegistrar)) {
	f(g.server)
}

func (g *Grpc) Close() {
	g.server.GracefulStop()
}
