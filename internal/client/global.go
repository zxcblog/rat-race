package client

import (
	"context"
	"github.com/zxcblog/rat-race/config"
	"github.com/zxcblog/rat-race/pkg/logger"
	"log"
)

var (
	// DB 数据库操作实例
	DB    *MariaDB
	Log   *logger.Logger
	Conf  *config.Config
	Redis *RedisDB
)

// Init 初始化全局信息
func Init(fileName string) error {
	var err error

	// 初始化配置信息
	if Conf, err = config.InitConfig(fileName); err != nil {
		log.Fatalf("配置文件初始化失败：%s", err.Error())
		return err
	}
	if err = SetConfig(Conf); err != nil {
		log.Fatalf("初始化配置项失败：%s", err.Error())
		return err
	}

	if DB, err = MariadbInit(DBConf); err != nil {
		log.Fatalf("Mariadb初始化失败：%s", err.Error())
		return err
	}

	if Redis, err = RedisInit(RedisConf); err != nil {
		log.Fatalf("Redis初始化失败：%s", err.Error())
		return err
	}

	Log = logger.NewLogger(LogConf)
	return nil
}

// Close 关闭全局变量
func Close() {
	ctx := context.Background()

	if err := DB.Close(); err != nil {
		Log.ErrorF(ctx, "关闭mariadb数据库失败:%s", err.Error())
	}

	if err := Log.Close(); err != nil {
		log.Fatalf("日志服务关闭失败：%s", err.Error())
	}
}
