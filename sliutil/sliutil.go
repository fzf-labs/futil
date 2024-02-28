//nolint:gosec
package sliutil

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// DropTop DropTop会返回剩下的切片在去掉上面n个元素后如果切片的元素少于n那么它会返回空切片如果n < 0它会返回空切片。
func DropTop[T any](ss []T, n int) (drop []T) {
	if n < 0 || n >= len(ss) {
		return
	}

	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #145.
	drop = make([]T, len(ss)-n)
	copy(drop, ss[n:])

	return
}

// FunctionValue return the reflect value of a function
func FunctionValue(function any) reflect.Value {
	v := reflect.ValueOf(function)
	if v.Kind() != reflect.Func {
		panic(fmt.Sprintf("Invalid function type, value of type %T", function))
	}
	return v
}

// CheckSliceCallbackFuncSignature Check func sign :  s :[]type1{} -> func(i int, data type1) type2
// see https://coolshell.cn/articles/21164.html#%E6%B3%9B%E5%9E%8BMap-Reduce
func CheckSliceCallbackFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	// Check it is a function
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	// first input param type should be int
	if fn.Type().In(0) != reflect.TypeOf(1) {
		return false
	}
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

// sliceElemType 获取切片元素类型
func sliceElemType(reflectType reflect.Type) reflect.Type {
	for {
		if reflectType.Kind() != reflect.Slice {
			return reflectType
		}

		reflectType = reflectType.Elem()
	}
}

var (
	bfPool = sync.Pool{
		New: func() any {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInt64ToString 将 int64 切片格式化为字符串，例如：n1,n2,n3。
func JoinInt64ToString(is []int64) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteByte(',')
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}

// SplitStringToInt64 将字符串拆分为 int64 切片。
func SplitStringToInt64(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}
	sArr := strings.Split(s, ",")
	res := make([]int64, 0, len(sArr))
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}
