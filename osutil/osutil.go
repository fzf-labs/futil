package osutil

import (
	"os"
	"runtime"
	"strconv"
)

// GetHostName 获取主机名
func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return hostname
}

// GetOS 获取操作系统名称
func GetOS() string {
	return runtime.GOOS
}

// IsWindows 检查当前操作系统是否为 Windows
func IsWindows() bool {
	return GetOS() == "windows"
}

// IsLinux 检查当前操作系统是否为 linux
func IsLinux() bool {
	return GetOS() == "linux"
}

// IsMac 检查当前操作系统是否为 macos
func IsMac() bool {
	return GetOS() == "darwin"
}

// GetOsEnv 通过键名获取环境变量的值。
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv 设置由键命名的环境变量的值。
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv 删除单个环境变量。
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// GetOsBits 获取此系统位 32 位或 64 位
// return bit int (32/64)
func GetOsBits() int {
	return 32 << (^uint(0) >> 63)
}

// GetGoroutineID 获得当前 goroutine id
func GetGoroutineID() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stk := string(buf[:n])
	str := stk[10:11]
	return strconv.Atoi(str)
}
