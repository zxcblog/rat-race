package client

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/zxcblog/rat-race/pkg/logger"
	"github.com/zxcblog/rat-race/pkg/metcd"
	"github.com/zxcblog/rat-race/pkg/tools"
	"go.uber.org/zap/zapcore"
)

var (
	Logger logger.ILogger
	Etcd   *metcd.MEtcd
	Config = new(Conf)

	Shutdown tools.ShutDowner
)

// Init 系统初始化方法,获取配置信息，初始化全局实例，
func Init(filename string) error {
	err := ReadConfig(filename)
	if err != nil {
		return err
	}

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

	// 全局关闭句柄
	Shutdown = tools.NewShutDown()

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
