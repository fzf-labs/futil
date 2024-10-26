package iputil

import "net"

// IsInternalIP 检测 IP 地址是否是内网地址
func IsInternalIP(ip string) bool {
	netIP := net.ParseIP(ip)
	if netIP.IsLoopback() {
		return true
	}
	ip4 := netIP.To4()
	if ip4 == nil {
		return false
	}
	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16 增加169字段
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
