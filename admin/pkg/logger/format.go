package logger

import (
	"context"
	"fmt"
)

func (log *Logger) InfoF(ctx context.Context, format string, v ...interface{}) {
	log.zapLogger.Info(fmt.Sprintf(format, v...))
}

func (log *Logger) ErrorF(ctx context.Context, format string, v ...interface{}) {
	log.zapLogger.Error(fmt.Sprintf(format, v...))
}

func (log *Logger) DebugF(ctx context.Context, format string, v ...interface{}) {
	log.zapLogger.Debug(fmt.Sprintf(format, v...))
}
