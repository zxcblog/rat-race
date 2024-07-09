package client

import (
	"github.com/zxcblog/rat-race/pkg/captcha"
	"github.com/zxcblog/rat-race/pkg/metcd"
	"github.com/zxcblog/rat-race/pkg/mgateway"
	"time"
)

type Conf struct {
	// 服务
	Server struct {
		Name     string
		LogLevel string
		GrpcConf struct {
			Port       string
			Host       string
			ServerName string
		}
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
	MariaDB struct {
		Host            string
		Port            string
		User            string
		Pass            string
		DbName          string
		MaxOpenConn     int
		ConnMaxLifeTime time.Duration
		MaxIdleConn     int
	}

	Redis struct {
		Host         string
		Port         string
		Pass         string
		Db           int
		MinIdleConns int
		PoolSize     int
	}

	Captcha captcha.Config
}
