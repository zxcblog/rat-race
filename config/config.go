package config

import (
	"github.com/zxcblog/rat-race/pkg/gateway"
	"github.com/zxcblog/rat-race/pkg/grpc"
	"github.com/zxcblog/rat-race/pkg/tools"
	"strings"
)

var (
	Server   *server
	GrpcConf *grpc.Config
	GwConf   *gateway.Config
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

	if err := setGwConf(conf); err != nil {
		return err
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
	err := conf.ReadConfig("Grpc", &c)
	if err != nil {
		return err
	}

	if GrpcConf == nil {
		GrpcConf = &grpc.Config{}
	}

	// 为热加载做预留
	GrpcConf.RunMode = Server.RunMode
	GrpcConf.Address = ":" + strings.Trim(c["port"], ":")
	if _, ok := c["transdatasize"]; ok {
		dataSize, err := tools.UnitConvInt64(c["transdatasize"], c["transdataunit"])
		if err != nil {
			return err
		}

		GrpcConf.TransDataSize = int(dataSize)
	}
	return nil
}

func setGwConf(conf *Config) error {
	c := make(map[string]string)
	err := conf.ReadConfig("Gateway", &c)
	if err != nil {
		return err
	}

	if GwConf == nil {
		GwConf = &gateway.Config{}
	}

	// 为热加载做预留
	GwConf.RunMode = Server.RunMode
	GwConf.Address = ":" + strings.Trim(c["port"], ":")
	GwConf.GrpcAddress = GrpcConf.Address
	GwConf.GrpcTransData = GrpcConf.TransDataSize
	return nil
}
