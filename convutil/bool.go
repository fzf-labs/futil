package conv

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

var (
	emptyStringMap = map[string]struct{}{
		"":      {},
		"0":     {},
		"no":    {},
		"off":   {},
		"false": {},
	}
)

// Bool converts `any` to bool.
func Bool(any any) (bool, error) {
	switch b := any.(type) {
	case nil:
		return false, nil
	case bool:
		return b, nil
	case int:
		return b != 0, nil
	case int64:
		return b != 0, nil
	case int32:
		return b != 0, nil
	case int16:
		return b != 0, nil
	case int8:
		return b != 0, nil
	case uint:
		return b != 0, nil
	case uint64:
		return b != 0, nil
	case uint32:
		return b != 0, nil
	case uint16:
		return b != 0, nil
	case uint8:
		return b != 0, nil
	case float64:
		return b != 0, nil
	case float32:
		return b != 0, nil
	case time.Duration:
		return b != 0, nil
	case []byte:
		if _, ok := emptyStringMap[strings.ToLower(string(b))]; ok {
			return false, nil
		}
		return true, nil
	case string:
		if _, ok := emptyStringMap[strings.ToLower(b)]; ok {
			return false, nil
		}
		return true, nil
	case json.Number:
		j, err := b.Int64()
		if err != nil {
			return false, err
		}
		return j != 0, nil
	case interface{ Bool() bool }:
		return b.Bool(), nil
	case interface{ IsZero() bool }:
		return !b.IsZero(), nil
	default:
		rv := reflect.ValueOf(any)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil(), nil
		case reflect.Map:
			return rv.Len() != 0, nil
		case reflect.Array:
			return rv.Len() != 0, nil
		case reflect.Slice:
			return rv.Len() != 0, nil
		case reflect.Struct:
			return true, nil
		default:
			return !rv.IsZero(), nil
		}
	}
}
