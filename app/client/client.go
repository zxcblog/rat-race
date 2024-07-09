package client

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/zxcblog/rat-race/pkg/captcha"
	"github.com/zxcblog/rat-race/pkg/logger"
	"github.com/zxcblog/rat-race/pkg/metcd"
	"github.com/zxcblog/rat-race/pkg/tools"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"time"
)

var (
	Logger  logger.ILogger
	Etcd    *metcd.MEtcd
	Config  = new(Conf)
	MariaDB *gorm.DB
	RedisDB *redis.Client

	Shutdown tools.ShutDowner
	Captcha  *base64Captcha.Captcha
)

// Init 系统初始化方法,获取配置信息，初始化全局实例，
func Init(filename string) error {
	err := ReadConfig(filename)
	if err != nil {
		return err
	}

	// 设置grpc和gateway启动时的服务名和监听的IP地址信息
	Config.Server.GrpcConf.Host = tools.GetSelfIP()
	Config.Server.GrpcConf.ServerName = Config.Server.Name

	// 日志初始化
	{
		Config.Log.FileLevel = Config.Server.LogLevel
		Config.Log.ConsoleLevel = Config.Server.LogLevel

		core := make([]zapcore.Core, 0, 2)
		if Config.Log.Filename != "" {
			core = append(core, logger.WithFileCore(logger.GetZapEncode(), Config.Log.Filename,
				Config.Log.MaxSize,
				Config.Log.MaxAges,
				Config.Log.MaxBackus,
				Config.Log.Compress,
				Config.Log.LocalTime,
				Config.Log.FileLevel,
			))
		} else {
			Config.Log.Console = true
		}

		if Config.Log.Console {
			core = append(core, logger.WithConsoleCore(logger.GetZapEncode(), Config.Log.ConsoleLevel))
		}
		Logger = logger.NewLogger(Config.Server.Name, core...)
	}

	// etcd 初始化
	Etcd, err = metcd.NewEtcd(Config.Etcd, Logger)
	if err != nil {
		return err
	}

	// gorm.DB 初始化
	MariaDB = MariadbInit()
	RedisDB = RedisInit()

	// 全局关闭句柄
	Shutdown = tools.NewShutDown()

	// 验证码配置
	Captcha, err = captcha.NewCaptcha(Config.Captcha, captcha.NewRedisStore(RedisDB, time.Duration(Config.Captcha.ExpireTime)*time.Minute, CaptchaRedisKey))
	if err != nil {
		Logger.ErrorF("验证码启动失败：%s", err.Error())
		return err
	}
	return nil
}

// ReadConfig 读取配置文件并写入Config全局变量中
func ReadConfig(filename string) error {
	path, err := tools.GetPath(filename)
	if err != nil {
		return err
	}

	if !tools.IsExists(path) {
		return errors.New("配置文件不存在")
	}

	vp := viper.New()
	vp.SetConfigFile(path)
	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	return vp.Unmarshal(&Config)
}
