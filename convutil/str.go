package conv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// String converts `any` to string.
func String(any any) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case nil:
		return ""
	case bool:
		return strconv.FormatBool(value)
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case []byte:
		return string(value)
	case string:
		return value
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	case json.Number:
		return value.String()
	case template.HTML:
		return string(value)
	case template.URL:
		return string(value)
	case template.JS:
		return string(value)
	case template.CSS:
		return string(value)
	case template.HTMLAttr:
		return string(value)
	case fmt.Stringer:
		return value.String()
	case error:
		return value.Error()
	default:
		// Reflect checks.
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if (kind == reflect.Chan || kind == reflect.Map || kind == reflect.Slice || kind == reflect.Func || kind == reflect.Ptr || kind == reflect.Interface || kind == reflect.UnsafePointer) && rv.IsNil() {
			return ""
		}
		if kind == reflect.String {
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		// Finally, we use json.Marshal to convert.
		jsonMarshal, err := json.Marshal(value)
		if err != nil {
			return fmt.Sprint(value)
		}
		return string(jsonMarshal)
	}
}

// Utf8ToGbk convert utf8 encoding data to GBK encoding data.
func Utf8ToGbk(bs []byte) ([]byte, error) {
	r := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewEncoder())
	b, err := io.ReadAll(r)
	return b, err
}

// GbkToUtf8 convert GBK encoding data to utf8 encoding data.
func GbkToUtf8(bs []byte) ([]byte, error) {
	r := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewDecoder())
	b, err := io.ReadAll(r)
	return b, err
}
