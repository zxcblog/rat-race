package logger

import (
	"context"
	"testing"
)

// 测试日志打印信息
func TestLogger(t *testing.T) {
	logger := NewLogger("rat-race", WithConsoleCore(nil, "info"))

	ctx := context.Background()

	logger.ErrorFWithCtx(ctx, "测试 error 级别日志:%s", "error")
	logger.InfoFWithCtx(ctx, "测试 info 级别日志:%s", "info")
}
