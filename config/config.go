package config

import (
	"github.com/zxcblog/rat-race/pkg/grpc"
	"strings"
)

var (
	Server   *server
	GrpcConf *grpc.Config
)

type server struct {
	Host    string
	RunMode string
}

func setConfig(conf *Config) error {
	if err := setServer(conf); err != nil {
		return nil
	}

	if err := setGrpcConf(conf); err != nil {
		return nil
	}

	return nil
}

func setServer(conf *Config) error {
	if Server == nil {
		Server = &server{}
		return conf.ReadConfig("Service", Server)
	}
	return nil
}

func setGrpcConf(conf *Config) error {
	c := make(map[string]string)
	if err := conf.ReadConfig("Grpc", &c); err != nil {
		return err
	}

	if GrpcConf == nil {
		GrpcConf = &grpc.Config{}
	}

	// 为热加载做预留
	GrpcConf.RunMode = Server.RunMode
	GrpcConf.Address = ":" + strings.Trim(c["port"], ":")
	return nil
}
