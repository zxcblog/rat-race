package metcd

import (
	"errors"
	"github.com/zxcblog/rat-race/framework/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

type MEtcdCfg struct {
	Endpoints   []string
	DialTimeOut int64
	Username    string
	Password    string
}

type MEtcd struct {
	DialTimeout int64

	log logger.ILogger
	cli *clientv3.Client

	LeasesMap map[string]*Lease
}

func NewEtcd(conf MEtcdCfg, log logger.ILogger) (*MEtcd, error) {
	dialTimeout := time.Duration(conf.DialTimeOut) * time.Second

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: dialTimeout,
		Username:    conf.Username,
		Password:    conf.Password,
	})
	if err != nil {
		log.DebugF("etcd连接失败：%s", err.Error())
		return nil, err
	}

	log.DebugF(`
etcd启动配置信息
	points  : %s
	root	: %s
	password: %s
`, strings.Join(conf.Endpoints, ","), conf.Username, conf.Password)
	return &MEtcd{
		cli:         cli,
		log:         log,
		DialTimeout: conf.DialTimeOut,
		LeasesMap:   make(map[string]*Lease),
	}, nil
}

// RegisterLease 注册租约
func (m *MEtcd) RegisterLease(key, val string) error {
	lease, err := newLease(m, key, val, m.DialTimeout)
	if err != nil {
		return err
	}

	m.LeasesMap[key] = lease
	go lease.keepAlive()

	return nil
}

// CloseLease 关闭租约
func (m *MEtcd) CloseLease(key string) error {
	lease, ok := m.LeasesMap[key]
	if !ok {
		return errors.New("错误的租约ID")
	}

	if err := lease.close(); err != nil {
		return err
	}
	delete(m.LeasesMap, key)
	return nil
}

// GetClient 获取etcd操作实例
func (m *MEtcd) GetClient() *clientv3.Client {
	return m.cli
}
