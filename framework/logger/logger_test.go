package logger

import (
	"context"
	"testing"
)

// 测试日志打印信息
func TestLogger(t *testing.T) {
	logger := NewLogger("rat-race", WithConsoleCore(nil, "info"))

	ctx := context.Background()

	logger.ErrorF(ctx, "测试 error 级别日志:%s", "error")
	logger.InfoF(ctx, "测试 info 级别日志:%s", "info")
}
