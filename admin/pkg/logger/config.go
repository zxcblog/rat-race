package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

const (
	// DefaultLevel the default log level
	DefaultLevel = zapcore.InfoLevel

	// DefaultTimeLayout the default time layout;
	DefaultTimeLayout = time.RFC3339
)

// AppendOut 多个日志记录输出目录及级别
type AppendOut struct {
}

type Config struct {
	MaxSize    int    // 保存大小
	MaxAge     int    // 保存时间
	MaxBackups int    // 备份保留数量
	Compress   bool   // 是否压缩
	LocalTime  bool   // 是否使用本地时间
	IsConsole  bool   // 是否输出到控制台
	FileName   string // 输出的文件
	Level      string // 日志等级
	Name       string // 日志服务名称
}

// GetLevel 设置日志输出大小
func (c *Config) GetLevel() zapcore.Level {
	// debug < info < warn < error
	switch c.Level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
