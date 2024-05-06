package logger

import (
	"context"
	"testing"
)

// 测试日志打印信息
func TestLogger(t *testing.T) {
	logger := NewLogger(&Config{
		MaxSize:    128,
		MaxAge:     30,
		MaxBackups: 300,
		Compress:   true,
		LocalTime:  true,
		IsConsole:  true,
		FileName:   "./log/error.log",
		Level:      "error",
		Name:       "rat-race",
	})

	ctx := context.Background()

	logger.ErrorF(ctx, "测试 error 级别日志:%s", "error")
	logger.InfoF(ctx, "测试 info 级别日志:%s", "info")
}
