package framework

import "github.com/zxcblog/rat-race/framework/logger"

type Options interface {
	apply(engine *Engine)
}

type OptionFunc func(*Engine)

func (o OptionFunc) apply(engine *Engine) {
	o(engine)
}

// WithLogInstance 设置日志实例
func WithLogInstance(log logger.ILogger) OptionFunc {
	return func(engine *Engine) {
		engine.log = log
	}
}
