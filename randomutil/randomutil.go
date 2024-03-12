package randomutil

import (
	"math/rand"
	"time"
)

// some consts string chars
const (
	Numbers       = "0123456789"
	AlphaNumLower = "abcdefghijklmnopqrstuvwxyz"
	AlphaNumUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandomSeed 随机数种子
func RandomSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt 在 [min, max) 处返回一个随机整数int
func RandomInt(min, max int) int {
	return min + RandomSeed().Intn(max-min)
}

// RandomInt64 在 [min, max) 处返回一个随机整数int64
func RandomInt64(min, max int64) int64 {
	return min + RandomSeed().Int63n(max-min)
}

// RandomStr 随机字符串
func RandomStr(n int) string {
	cs := make([]byte, n)
	str := Numbers + AlphaNumLower + AlphaNumUpper
	sl := len(str)
	for i := 0; i < n; i++ {
		// 1607400451937462000
		idx := RandomSeed().Intn(sl) // 0 - 25
		cs[i] = str[idx]
	}
	return string(cs)
}

// RandomChars 随机字符串
func RandomChars(n int, char ...string) string {
	cs := make([]byte, n)
	str := ""
	if len(char) > 0 {
		for _, s := range char {
			str += s
		}
	} else {
		str = Numbers + AlphaNumLower + AlphaNumUpper
	}
	sl := len(str)
	for i := 0; i < n; i++ {
		// 1607400451937462000
		idx := RandomSeed().Intn(sl) // 0 - 25
		cs[i] = str[idx]
	}
	return string(cs)
}

// RandomNumber 随机生成指定长度的数字
func RandomNumber(n int) string {
	cs := make([]byte, n)
	str := Numbers
	sl := len(str)
	for i := 0; i < n; i++ {
		// 1607400451937462000
		idx := RandomSeed().Intn(sl) // 0 - 25
		cs[i] = str[idx]
	}
	return string(cs)
}
