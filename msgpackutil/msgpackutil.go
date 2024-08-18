package msgpackutil

import (
	"github.com/vmihailenco/msgpack/v5"
)

// Marshal 编码
func Marshal(v any) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal 解码
func Unmarshal(bts []byte, ptr any) error {
	return msgpack.Unmarshal(bts, ptr)
}

// MarshalString 编码到字符串
func MarshalString(v any) (string, error) {
	marshal, err := msgpack.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

// UnmarshalString 解码字符串
func UnmarshalString(str string, ptr any) error {
	return msgpack.Unmarshal([]byte(str), ptr)
}
