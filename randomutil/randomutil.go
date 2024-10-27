package randomutil

import (
	"math/rand/v2"
	"strings"
)

// some consts string chars
const (
	Numbers       = "0123456789"
	AlphaNumLower = "abcdefghijklmnopqrstuvwxyz"
	AlphaNumUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandomInt 在 [min, max) 处返回一个随机整数int
func RandomInt(minItem, maxItem int) int {
	if minItem >= maxItem {
		panic("minItem must be less than maxItem")
	}
	return minItem + rand.IntN(maxItem-minItem)
}

// RandomInt64 在 [min, max) 处返回一个随机整数int64
func RandomInt64(minItem, maxItem int64) int64 {
	if minItem >= maxItem {
		panic("minItem must be less than maxItem")
	}
	return minItem + rand.Int64N(maxItem-minItem)
}

// RandomStr 随机字符串
func RandomStr(n int) string {
	cs := make([]byte, n)
	str := Numbers + AlphaNumLower + AlphaNumUpper
	sl := len(str)
	for i := 0; i < n; i++ {
		idx := rand.IntN(sl)
		cs[i] = str[idx]
	}
	return string(cs)
}

// RandomChars 随机字符串
func RandomChars(n int, char ...string) string {
	if n <= 0 {
		return ""
	}
	var builder strings.Builder
	if len(char) > 0 {
		for _, s := range char {
			builder.WriteString(s)
		}
	} else {
		builder.WriteString(Numbers + AlphaNumLower + AlphaNumUpper)
	}
	str := builder.String()
	sl := len(str)
	for i := 0; i < n; i++ {
		idx := rand.IntN(sl)
		cs := str[idx]
		builder.WriteByte(cs)
	}
	return builder.String()
}

// RandomNumber 随机生成指定长度的数字
func RandomNumber(n int) string {
	cs := make([]byte, n)
	str := Numbers
	sl := len(str)
	for i := 0; i < n; i++ {
		idx := rand.IntN(sl)
		cs[i] = str[idx]
	}
	return string(cs)
}
