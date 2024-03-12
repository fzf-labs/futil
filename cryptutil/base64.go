package cryptutil

import (
	"github.com/golang-module/dongle"
)

// Base64Encode 使用 base64 编码对字符串进行编码
func Base64Encode(s string) string {
	return dongle.Encode.FromString(s).ByBase64().ToString()
}

// Base64Decode 解码 base64 编码的字符串
func Base64Decode(s string) string {
	return dongle.Decode.FromString(s).ByBase64().ToString()
}
