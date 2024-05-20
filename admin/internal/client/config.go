package client

import (
	"github.com/zxcblog/rat-race/config"
	"github.com/zxcblog/rat-race/pkg/captcha"
	"github.com/zxcblog/rat-race/pkg/gateway"
	"github.com/zxcblog/rat-race/pkg/grpc"
	"github.com/zxcblog/rat-race/pkg/logger"
	"github.com/zxcblog/rat-race/pkg/tools"
	"strings"
)

const (
	RUN_MODE_DEV     = "dev"     // dev
	RUN_MODE_RELEASE = "release" // 正式环境
)

var (
	Server      *server
	GrpcConf    *grpc.Config
	GwConf      *gateway.Config
	DBConf      *dbConfig
	LogConf     *logger.Config
	RedisConf   *redisConfig
	CaptchaConf *captcha.Config
)

type server struct {
	Host    string
	RunMode string
	Name    string
}

func SetConfig(conf *config.Config) error {
	if err := setServer(conf); err != nil {
		return nil
	}

	if err := setGrpcConf(conf); err != nil {
		return nil
	}

	if err := setGwConf(conf); err != nil {
		return err
	}

	if err := setDBConf(conf); err != nil {
		return err
	}

	if err := setLogConf(conf); err != nil {
		return err
	}

	if err := setRedisConf(conf); err != nil {
		return err
	}

	if err := setCaptchaConf(conf); err != nil {
		return err
	}
	return nil
}

// 设置服务启动的基本配置
func setServer(conf *config.Config) error {
	if Server == nil {
		Server = &server{}
		return conf.ReadConfig("Service", Server)
	}
	return nil
}

// 设置grpc的基本配置
func setGrpcConf(conf *config.Config) error {
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

// 设置gateway基本配置
func setGwConf(conf *config.Config) error {
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

func setDBConf(conf *config.Config) error {
	return conf.ReadConfig("Mariadb", &DBConf)
}

func setLogConf(conf *config.Config) error {
	if err := conf.ReadConfig("Logger", &LogConf); err != nil {
		return err
	}
	LogConf.Name = Server.Name
	return nil
}

func setRedisConf(conf *config.Config) error {
	return conf.ReadConfig("Redis", &RedisConf)
}

func setCaptchaConf(conf *config.Config) error {
	return conf.ReadConfig("Captcha", &CaptchaConf)
}
