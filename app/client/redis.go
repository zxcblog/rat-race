package client

//
//import (
//	"context"
//	"fmt"
//	"github.com/redis/go-redis/v9"
//	"github.com/zxcblog/rat-race/pkg/starter"
//	"github.com/zxcblog/rat-race/pkg/tools"
//)
//
//type redisConfig struct {
//	Host         string
//	Port         string
//	Pass         string
//	Db           int
//	MinIdleConns int
//	PoolSize     int
//}
//
//type redisDB struct {
//	*redis.Client
//}
//
//// RedisInit 数据库初始化
//func RedisInit(conf *redisConfig) (*redisDB, error) {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:         fmt.Sprintf("%s:%s", conf.Host, conf.Port),
//		Password:     conf.Pass,
//		DB:           conf.Db,
//		MinIdleConns: conf.MinIdleConns,
//		PoolSize:     conf.PoolSize,
//	})
//
//	if err := rdb.Ping(context.Background()).Err(); err != nil {
//		return nil, err
//	}
//
//	redisClient := &redisDB{rdb}
//	redisClient.registerComp(conf)
//	return redisClient, nil
//}
//
//// Close 关闭数据库连接
//func (db *redisDB) Close() error {
//	return db.Client.Close()
//}
//
//// registerComp
//func (db *redisDB) registerComp(conf *redisConfig) {
//	// 每次启动都打印
//	comp := starter.NewComp("Redis", true)
//
//	comp.SetCompItem("pass", conf.Pass)
//	comp.SetCompItem("host", conf.Host)
//	comp.SetCompItem("port", conf.Port)
//	comp.SetCompItem("db", tools.TransStr(conf.Db))
//}
