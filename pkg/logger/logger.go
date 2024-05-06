package logger

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://dbwu.tech/posts/golang_zap/
type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger(config *Config) *Logger {
	// 创建基本日志包信息
	encod := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "file",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,        // 定义日志默认行结尾
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // 大写并添加颜色的字符串
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
		}, // 2006-01-02 15:04:05.000 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,

		EncodeName: zapcore.FullNameEncoder,
		//EncodeName: func(s string, encoder zapcore.PrimitiveArrayEncoder) {
		//	encoder.AppendString(s)
		//},
		EncodeCaller:     zapcore.ShortCallerEncoder, // 路径
		ConsoleSeparator: "  ",

		// EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		// 	enc.AppendInt64(int64(d) / 1000000)
		// },
	}

	cores := make([]zapcore.Core, 0, 2)
	// 是否写入到控制台
	if config.IsConsole {
		stdout := zapcore.Lock(os.Stdout) //可能会有多个日志打印，使用互斥锁，输出到控制台
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encod), zapcore.NewMultiWriteSyncer(stdout), config.GetLevel()))
	}

	// 写入到文件中
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		Compress:   config.Compress,
		LocalTime:  config.LocalTime,
	})
	encod.EncodeLevel = zapcore.CapitalLevelEncoder
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encod), writer, config.GetLevel()))

	// 创建一个多路复用的日志核心（zapcore.Core）,同时将日志输出到多个目的地，该函数接收任意数量的zapcore.Core参数
	// 并返回一个心得zapcore.Core实例。 当写入日志时，会将相同得日志条目发送给所有传入得核心实例
	// zap.Config 中的 OutputPaths 也可以实现相同的功能， 但不能直接决定日志是否会被同时写入多个地方
	// NewTee 创建的Core写入Logger时， 会被复制并分别写入Tee中包含的Core所指向的输出位置
	zapLogger := zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
	if len(config.Name) > 0 {
		zapLogger = zapLogger.Named(config.Name) // 输出服务名称
	}

	return &Logger{
		zapLogger: zapLogger,
	}
}
