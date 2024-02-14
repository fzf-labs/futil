package conv

import (
	"fmt"
	"strconv"
)

// Uint converts `any` to uint.
func Uint(any any) (uint, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(uint); ok {
		return v, nil
	}
	u, err := Uint64(any)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}

// Uint8 converts `any` to uint8.
func Uint8(any any) (uint8, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(uint8); ok {
		return v, nil
	}
	u, err := Uint64(any)
	if err != nil {
		return 0, err
	}
	return uint8(u), nil
}

// Uint16 converts `any` to uint16.
func Uint16(any any) (uint16, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(uint16); ok {
		return v, nil
	}
	u, err := Uint64(any)
	if err != nil {
		return 0, err
	}
	return uint16(u), nil
}

// Uint32 converts `any` to uint32.
func Uint32(any any) (uint32, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(uint32); ok {
		return v, nil
	}
	u, err := Uint64(any)
	if err != nil {
		return 0, err
	}
	return uint32(u), nil
}

// Uint64 converts `any` to uint64.
func Uint64(any any) (uint64, error) {
	if any == nil {
		return 0, nil
	}
	switch value := any.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return uint64(value), nil
	case int8:
		return uint64(value), nil
	case int16:
		return uint64(value), nil
	case int32:
		return uint64(value), nil
	case int64:
		return uint64(value), nil
	case uint:
		return uint64(value), nil
	case uint8:
		return uint64(value), nil
	case uint16:
		return uint64(value), nil
	case uint32:
		return uint64(value), nil
	case uint64:
		return value, nil
	case float32:
		return uint64(value), nil
	case float64:
		return uint64(value), nil
	case []byte:
		return strconv.ParseUint(string(value), 10, 64)
	case string:
		return strconv.ParseUint(value, 10, 64)
	default:
		return 0, fmt.Errorf("unable to conv %#v of type %T to int64", any, any)
	}
}
