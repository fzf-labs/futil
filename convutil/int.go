package conv

import (
	"fmt"
	"strconv"
)

// Int converts `any` to int.
func Int(any any) (int, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(int); ok {
		return v, nil
	}
	i, err := Int64(any)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// Int8 converts `any` to int8.
func Int8(any any) (int8, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(int8); ok {
		return v, nil
	}
	i, err := Int64(any)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

// Int16 converts `any` to int16.
func Int16(any any) (int16, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(int16); ok {
		return v, nil
	}
	i, err := Int64(any)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

// Int32 converts `any` to int32.
func Int32(any any) (int32, error) {
	if any == nil {
		return 0, nil
	}
	if v, ok := any.(int32); ok {
		return v, nil
	}
	i, err := Int64(any)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// Int64 converts `any` to int64.
func Int64(any any) (int64, error) {
	if any == nil {
		return 0, nil
	}
	switch value := any.(type) {
	case int:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int64:
		return value, nil
	case uint:
		return int64(value), nil
	case uint8:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case uint64:
		return int64(value), nil
	case float32:
		return int64(value), nil
	case float64:
		return int64(value), nil
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return strconv.ParseInt(string(value), 10, 64)
	case string:
		return strconv.ParseInt(value, 10, 64)
	default:
		return 0, fmt.Errorf("unable to conv %#v of type %T to int64", any, any)
	}
}
