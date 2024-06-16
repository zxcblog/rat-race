package tools

import (
	"net"
	"runtime"
)

func GetSelfIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	if runtime.GOOS == "windows" { // 只有windows机器有这个网卡问题
		return "127.0.0.1"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
