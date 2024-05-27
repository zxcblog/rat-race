package framework

import (
	"github.com/zxcblog/rat-race/framework/gateway"
	"github.com/zxcblog/rat-race/framework/grpc"
	"github.com/zxcblog/rat-race/framework/logger"
)

// Engine 框架的核心，用来启动 grpc 和 grpc-gateway 服务
type Engine struct {
	*grpc.Grpc

	*gateway.Gateway

	log logger.ILogger
}

// New 初始化框架信息
func New(options ...Options) *Engine {
	engine := &Engine{
		log: logger.NewLogger("framework"),
	}

	for _, o := range options {
		o.apply(engine)
	}

	// 添加http服务
	engine.Gateway = gateway.New(engine.log)

	engine.log.DebugF("服务初始化完成，正在启动中")
	return engine
}

func (e *Engine) Run() {
	e.Grpc.Run()
	e.Gateway.Run()
	e.log.DebugF("服务启动完毕")
}
