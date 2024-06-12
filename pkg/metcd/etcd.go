package metcd

import (
	"github.com/zxcblog/rat-race/framework/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type MEtcdCfg struct {
	Endpoints   []string
	DialTimeOut int
	Username    string
	Password    string
}

type MEtcd struct {
	log logger.ILogger
	cli *clientv3.Client
}

func NewEtcd(conf MEtcdCfg, log logger.ILogger) (*MEtcd, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: time.Duration(conf.DialTimeOut) * time.Minute,
		Username:    conf.Username,
		Password:    conf.Password,
	})
	if err != nil {
		log.DebugF("etcd连接失败：%s", err.Error())
		return nil, err
	}

	return &MEtcd{cli: cli, log: log}, nil
}
