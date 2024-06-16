package tools

import "testing"

func TestGetSelfIP(t *testing.T) {
	t.Error(GetSelfIP())
	//addrs, err := net.InterfaceAddrs()
	//if err != nil {
	//	return ""
	//}
	//if env.IsDevMod() && runtime.GOOS == "windows" { // 只有windows机器有这个网卡问题
	//	return "127.0.0.1"
	//}
	//for _, address := range addrs {
	//	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//		if ipnet.IP.To4() != nil {
	//			return ipnet.IP.String()
	//		}
	//	}
	//}
	//return ""
}
