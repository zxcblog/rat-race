package gateway

import "time"

const (
	DevMod     = "dev"     // dev
	ReleaseMod = "release" // 正式环境
)

// Config 启动配置信息
type Config struct {
	Address       string
	RunMode       string
	GrpcAddress   string
	GrpcTransData int
	ShutDownTime  time.Duration
}
