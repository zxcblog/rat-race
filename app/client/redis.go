package client

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// RedisInit 数据库初始化
func RedisInit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", Config.Redis.Host, Config.Redis.Port),
		Password:     Config.Redis.Pass,
		DB:           Config.Redis.Db,
		MinIdleConns: Config.Redis.MinIdleConns,
		PoolSize:     Config.Redis.PoolSize,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		Logger.ErrorF("redis连接失败：%s", err.Error())
		panic(err.Error())
	}

	return rdb
}

//// Close 关闭数据库连接
//func (db *redisDB) Close() error {
//	return db.Client.Close()
//}
