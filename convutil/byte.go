package conv

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math"
	"reflect"
	"strconv"
)

// Byte converts `any` to byte.
func Byte(any any) (byte, error) {
	if v, ok := any.(byte); ok {
		return v, nil
	}
	return Uint8(any)
}

// Bytes converts `any` to []byte.
func Bytes(any any) ([]byte, error) {
	if any == nil {
		return nil, nil
	}
	v := reflect.ValueOf(any)
	switch any.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, bits)
		return b, nil
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, bits)
		return b, nil
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	default:
		newValue, err := json.Marshal(any)
		return newValue, err
	}
}
