package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// GetLevel 设置日志输出大小
func GetLevel(level string) zapcore.Level {
	// debug < info < warn < error
	switch level {
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

// GetZapEncode 获取到默认的日志配置信息
func GetZapEncode() *zapcore.EncoderConfig {
	return &zapcore.EncoderConfig{
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
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder, // 路径
		ConsoleSeparator: "  ",
	}
}

// WithFileCore 添加文件核心Core
// filename 写入的文件名称, maxSize 文件大小, maxAges 文件保存日志, maxBackus 最大保存数量 int, isCompress 是否压缩, isLocalTime 是否使用本地时间 bool
func WithFileCore(encode *zapcore.EncoderConfig, filename string, maxSize, maxAges, maxBackus int, isCompress, isLocalTime bool, logLevel string) zapcore.Core {
	if encode == nil {
		encode = GetZapEncode()
		// 写入到文件中的日志， 日志级别使用不带颜色大写字母
		encode.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAges,
		MaxBackups: maxBackus,
		LocalTime:  isLocalTime,
		Compress:   isCompress,
	})

	return zapcore.NewCore(zapcore.NewJSONEncoder(*encode), writer, GetLevel(logLevel))
}

// WithConsoleCore 日志输出到控制台
func WithConsoleCore(encode *zapcore.EncoderConfig, logLevel string) zapcore.Core {
	if encode == nil {
		encode = GetZapEncode()
	}

	stdout := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(zapcore.NewConsoleEncoder(*encode), zapcore.NewMultiWriteSyncer(stdout), GetLevel(logLevel))
}
