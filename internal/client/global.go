package client

import (
	"github.com/zxcblog/rat-race/pkg/logger"
)

var (
	// DB 数据库操作实例
	DB  *MariaDB
	Log *logger.Logger
)
