package tools

import (
	"os"
	"os/signal"
	"syscall"
)

// ShutDowner 优雅关闭服务的接口,
type ShutDowner interface {
	// WithSignals 监听更多的信号
	WithSignals(signals ...syscall.Signal) ShutDowner

	// Close 注册关闭服务
	Close(funcs ...func())
}

type ShutDown struct {
	ctx chan os.Signal
}

// NewShutDown 创建服务, 默认监听SIGINT和SIGTERM
func NewShutDown() ShutDowner {
	shutDown := &ShutDown{
		ctx: make(chan os.Signal, 1),
	}

	return shutDown.WithSignals(syscall.SIGINT, syscall.SIGTERM)
}

func (s *ShutDown) WithSignals(signals ...syscall.Signal) ShutDowner {
	for _, sig := range signals {
		signal.Notify(s.ctx, sig)
	}
	return s
}

func (s *ShutDown) Close(funcs ...func()) {
	<-s.ctx
	signal.Stop(s.ctx)

	for _, f := range funcs {
		f()
	}
}
