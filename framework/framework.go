package framework

import (
	"github.com/zxcblog/rat-race/framework/gateway"
)

// Engine 框架的核心，用来启动 grpc 和 grpc-gateway 服务
type Engine struct {
	*gateway.Gateway
}

// New 初始化框架信息
func New() *Engine {
	return &Engine{
		Gateway: gateway.New(),
	}
}

func (e *Engine) Run() {
	e.Gateway.Run()
}
