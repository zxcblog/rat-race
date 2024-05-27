package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ILogger 日志接口信息
type ILogger interface {
	Close() error
	InfoF(format string, v ...interface{})
	InfoFWithCtx(ctx context.Context, format string, v ...interface{})
	ErrorF(format string, v ...interface{})
	ErrorFWithCtx(ctx context.Context, format string, v ...interface{})
	DebugF(format string, v ...interface{})
	DebugFWithCtx(ctx context.Context, format string, v ...interface{})
}

// https://dbwu.tech/posts/golang_zap/
type Logger struct {
	zapLogger *zap.Logger
	name      string
}

func NewLogger(name string, cores ...zapcore.Core) *Logger {
	// 默认使用控制台进行输出
	if len(cores) < 1 {
		cores = []zapcore.Core{WithConsoleCore(nil, "info")}
	}

	// 创建一个多路复用的日志核心（zapcore.Core）,同时将日志输出到多个目的地，该函数接收任意数量的zapcore.Core参数
	// 并返回一个心得zapcore.Core实例。 当写入日志时，会将相同得日志条目发送给所有传入得核心实例
	// zap.Config 中的 OutputPaths 也可以实现相同的功能， 但不能直接决定日志是否会被同时写入多个地方
	// NewTee 创建的Core写入Logger时， 会被复制并分别写入Tee中包含的Core所指向的输出位置
	zapLogger := zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
	if len(name) > 0 {
		zapLogger = zapLogger.Named(name) // 输出服务名称
	}

	return &Logger{
		zapLogger: zapLogger,
	}
}

// Close 关闭日志服务
func (l *Logger) Close() error {
	return l.zapLogger.Sync()
}

func (log *Logger) InfoF(format string, v ...interface{}) {
	log.zapLogger.Info(fmt.Sprintf(format, v...))
}

func (log *Logger) ErrorF(format string, v ...interface{}) {
	log.zapLogger.Error(fmt.Sprintf(format, v...))
}

func (log *Logger) DebugF(format string, v ...interface{}) {
	log.zapLogger.Debug(fmt.Sprintf(format, v...))
}

func (log *Logger) InfoFWithCtx(ctx context.Context, format string, v ...interface{}) {
	log.zapLogger.Info(fmt.Sprintf(format, v...))
}

func (log *Logger) ErrorFWithCtx(ctx context.Context, format string, v ...interface{}) {
	log.zapLogger.Error(fmt.Sprintf(format, v...))
}

func (log *Logger) DebugFWithCtx(ctx context.Context, format string, v ...interface{}) {
	log.zapLogger.Debug(fmt.Sprintf(format, v...))
}
