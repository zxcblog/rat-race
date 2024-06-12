package client

import (
	"github.com/zxcblog/rat-race/pkg/metcd"
	"github.com/zxcblog/rat-race/pkg/mgateway"
	"github.com/zxcblog/rat-race/pkg/mgrpc"
)

type Conf struct {
	// 服务
	Server struct {
		Name        string
		RunMode     string
		GrpcConf    mgrpc.GrpcConf
		GatewayConf mgateway.GatewayConf
	}
	//etcd配置信息
	Etcd metcd.MEtcdCfg
	//	Log 日志文件输出信息
	Log struct {
		MaxSize      int
		MaxAges      int
		MaxBackus    int
		Compress     bool
		LocalTime    bool
		Console      bool
		Filename     string
		FileLevel    string
		ConsoleLevel string
	}
}
