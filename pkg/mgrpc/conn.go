package mgrpc

import (
	"fmt"
	"github.com/zxcblog/rat-race/framework/logger"
	"github.com/zxcblog/rat-race/pkg/metcd"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetConnByEtcd 通过ETCD对服务进行发现并获取注册信息
func GetConnByEtcd(serverName string, etcd *metcd.MEtcd, log logger.ILogger) *grpc.ClientConn {
	etcdResolver, err := resolver.NewBuilder(etcd.GetClient())
	if err != nil {
		log.ErrorF("读取etcd-endpoint异常，程序强制退出：%s", err.Error())
		panic(err.Error())
	}

	conn, err := grpc.NewClient(fmt.Sprintf("etcd:///%s/%s", ServerRegister, serverName),
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.ErrorF("通过etcd对服务进行发现失败：%s", err.Error())
	}

	return conn

}
