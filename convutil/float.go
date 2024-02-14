package conv

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Float32 converts `any` to float32.
func Float32(any any) (float32, error) {
	switch s := any.(type) {
	case nil:
		return 0, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int:
		return float32(s), nil
	case int8:
		return float32(s), nil
	case int16:
		return float32(s), nil
	case int32:
		return float32(s), nil
	case int64:
		return float32(s), nil
	case float32:
		return s, nil
	case float64:
		return float32(s), nil
	case uint:
		return float32(s), nil
	case uint8:
		return float32(s), nil
	case uint16:
		return float32(s), nil
	case uint32:
		return float32(s), nil
	case uint64:
		return float32(s), nil
	case []byte:
		v, err := strconv.ParseFloat(string(s), 32)
		if err != nil {
			return 0, err
		}
		return float32(v), nil
	case string:
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, err
		}
		return float32(v), nil
	case json.Number:
		v, err := s.Float64()
		if err != nil {
			return 0, err
		}
		return float32(v), nil
	default:
		return 0, fmt.Errorf("unable to conv %#v of type %T to float32", any, any)
	}
}

// Float64 converts `any` to float64.
func Float64(any interface{}) (float64, error) {
	switch s := any.(type) {
	case nil:
		return 0, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case float32:
		return float64(s), nil
	case float64:
		return s, nil
	case uint:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case []byte:
		v, err := strconv.ParseFloat(string(s), 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	case json.Number:
		v, err := s.Float64()
		if err != nil {
			return 0, err
		}
		return v, nil
	default:
		return 0, fmt.Errorf("unable to conv %#v of type %T to float64", any, any)
	}
}
