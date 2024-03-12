package iputil

import (
	"context"
	"encoding/json"
	"io"
	"math"
	"net"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// GetLocalIPList 内网IP
func GetLocalIPList() (ips []string, err error) {
	ips = make([]string, 0)
	ices, e := net.Interfaces()
	if e != nil {
		return ips, e
	}
	for _, ice := range ices {
		if ice.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if ice.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		// ignore docker and warden bridge
		if strings.HasPrefix(ice.Name, "docker") || strings.HasPrefix(ice.Name, "w-") {
			continue
		}
		adders, e := ice.Addrs()
		if e != nil {
			return ips, e
		}
		for _, addr := range adders {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			ipStr := ip.String()
			if IsInternalIP(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}
	return ips, nil
}

// GetLocalIP 获取本机内网IP
func GetLocalIP() string {
	localIP := "127.0.0.1"
	ips, _ := GetLocalIPList()
	if len(ips) > 0 {
		localIP = ips[0]
	}
	return localIP
}

// GetPublicIP 获取公网ip
func GetPublicIP() string {
	conn, err := net.Dial("udp", "114.114.114.114:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// GetPublicIPByHTTP 获取公网ip
func GetPublicIPByHTTP() (string, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://httpbin.org/ip", http.NoBody)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type IPInfo struct {
		Origin string `json:"origin"`
	}
	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		return "", err
	}
	return ipInfo.Origin, nil
}

// GetRealIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func GetRealIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !IsInternalIP(ip) {
			return ip
		}
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-IP"))
	if ip != "" && !IsInternalIP(ip) {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !IsInternalIP(ip) {
			return ip
		}
	}
	return ""
}

// IPToLong 把ip字符串转为数值
func IPToLong(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}
	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// LongToIP 把数值转为ip字符串
func LongToIP(long uint) (string, error) {
	if long > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}
	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(long >> 24)
	ip[1] = byte(long >> 16)
	ip[2] = byte(long >> 8)
	ip[3] = byte(long)
	return ip.String(), nil
}
