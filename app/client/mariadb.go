package client

import (
	"fmt"
	"github.com/zxcblog/rat-race/app/model/mariadb"
	"gorm.io/gorm/logger"
	//"github.com/zxcblog/rat-race/internal/model/mariadb"
	//"github.com/zxcblog/rat-race/pkg/starter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// MariadbInit 数据库初始化
func MariadbInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		Config.MariaDB.User,
		Config.MariaDB.Pass,
		Config.MariaDB.Host,
		Config.MariaDB.Port,
		Config.MariaDB.DbName,
		true,
		"Local")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})
	if err != nil {
		Logger.ErrorF("[数据库连接失败] 数据库名称：%s, 错误信息：%s", Config.MariaDB.DbName, err.Error())
		panic(err.Error())
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	sqlDB, err := db.DB()

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(Config.MariaDB.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(Config.MariaDB.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * Config.MariaDB.ConnMaxLifeTime)

	logLevel := logger.Error
	if Config.Server.LogLevel == "debug" {
		logLevel = logger.Info
	}

	db.Session(&gorm.Session{Logger: db.Logger.LogMode(logLevel)})

	Logger.DebugF(`
mariadb 启动配置信息 dsn : %s`, dsn)
	if err = mariadb.InitSql(db); err != nil {
		Logger.ErrorF("数据迁移失败：%s", err.Error())
		panic(err.Error())
	}
	return db
}

//// Close 关闭数据库连接
//func (db *mariaDB) Close() error {
//	sqldb, err := db.DB.DB()
//	if err != nil {
//		return err
//	}
//	return sqldb.Close()
//}
