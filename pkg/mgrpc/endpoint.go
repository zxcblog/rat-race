package mgrpc

import (
	"encoding/json"
	"strings"
)

// ServerRegister 服务注册与发现的前缀
const ServerRegister = "/ms-endpoint"

type Endpoint struct {
	Addr string `json:"addr"`
	Pkgs string `json:"pkgs"`
}

// 将本项目实现的包信息注册到etcd环境中
func (g *Grpc) register(pkgs []string) error {
	endpoint := Endpoint{
		Addr: g.config.Addr(),
		Pkgs: strings.Join(pkgs, ","),
	}

	data, _ := json.Marshal(endpoint)

	return g.metcd.RegisterLease(strings.Join([]string{ServerRegister, g.config.ServerName, endpoint.Addr}, "/"), string(data))
}
