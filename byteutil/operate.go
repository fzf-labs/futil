package byteutil

import "unsafe"

// BytesToString 字节切片转字符串
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
