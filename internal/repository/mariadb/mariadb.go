package mariadb

import "gorm.io/gorm"

type DBConf struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

type DB struct {
	DB   *gorm.DB
	Conf *DBConf
}

func Conn(conf *DBConf) *DB {

	return &DB{
		Conf: conf,
	}

}
