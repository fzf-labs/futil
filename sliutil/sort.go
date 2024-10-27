package sliutil

import (
	"crypto/rand"
	"encoding/binary"
	"time"

	"golang.org/x/exp/constraints"
)

// Shuffle the slice using crypto/rand for better randomization
func Shuffle[T any](collection []T) []T {
	n := len(collection)
	if n <= 1 {
		return collection
	}

	for i := n - 1; i > 0; i-- {
		var buf [8]byte
		_, err := rand.Read(buf[:])
		if err != nil {
			// 如果加密随机数生成失败，回退到时间种子
			j := int(time.Now().UnixNano() % int64(i+1))
			collection[i], collection[j] = collection[j], collection[i]
			continue
		}
		// Safely convert to avoid integer overflow
		j := int(binary.BigEndian.Uint64(buf[:]) % uint64(uint(i+1)))
		collection[i], collection[j] = collection[j], collection[i]
	}
	return collection
}

// IsAsc 检查切片是否按升序排列。
func IsAsc[T constraints.Ordered](collection []T) bool {
	for i := 1; i < len(collection); i++ {
		if collection[i-1] > collection[i] {
			return false
		}
	}
	return true
}

// IsDesc 检查切片是否按降序排列。
func IsDesc[T constraints.Ordered](collection []T) bool {
	for i := 1; i < len(collection); i++ {
		if collection[i-1] < collection[i] {
			return false
		}
	}
	return true
}

// IsSorted 检查切片是否排序(升序或降序)。
func IsSorted[T constraints.Ordered](collection []T) bool {
	return IsAsc(collection) || IsDesc(collection)
}

// IsSortedByKey 检查切片是否按迭代函数排序。
func IsSortedByKey[T any, K constraints.Ordered](collection []T, iteratee func(item T) K) bool {
	size := len(collection)
	isAsc := func(data []T) bool {
		for i := 0; i < size-1; i++ {
			if iteratee(data[i]) > iteratee(data[i+1]) {
				return false
			}
		}
		return true
	}
	isDesc := func(data []T) bool {
		for i := 0; i < size-1; i++ {
			if iteratee(data[i]) < iteratee(data[i+1]) {
				return false
			}
		}
		return true
	}
	return isAsc(collection) || isDesc(collection)
}
