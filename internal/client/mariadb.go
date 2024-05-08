package client

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DBConf struct {
	Host            string
	Port            string
	User            string
	Pass            string
	DbName          string
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
	MaxIdleConn     int
}

type MariaDB struct {
	DB   *gorm.DB
	Conf *DBConf
}

func Conn(conf *DBConf) (*MariaDB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.DbName,
		true,
		"Local")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[数据库连接失败] 数据库名称：%s", conf.DbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	sqlDB, err := db.DB()

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(conf.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(conf.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * conf.ConnMaxLifeTime)

	return &MariaDB{DB: db, Conf: conf}, nil
}

// Close 关闭数据库连接
func (db *MariaDB) Close() error {
	sqldb, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqldb.Close()
}
