package sliutil

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

// Shuffle the slice.
// Play: https://go.dev/play/p/YHvhnWGU3Ge
func Shuffle[T any](slice []T) []T {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

// IsAsc 检查切片是否按升序排列。
func IsAsc[T constraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

// IsDesc 检查切片是否按降序排列。
func IsDesc[T constraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			return false
		}
	}
	return true
}

// IsSorted 检查切片是否排序(升序或降序)。
func IsSorted[T constraints.Ordered](slice []T) bool {
	return IsAsc(slice) || IsDesc(slice)
}

// IsSortedByKey 检查切片是否按迭代函数排序。
func IsSortedByKey[T any, K constraints.Ordered](slice []T, iteratee func(item T) K) bool {
	size := len(slice)
	isAscending := func(data []T) bool {
		for i := 0; i < size-1; i++ {
			if iteratee(data[i]) > iteratee(data[i+1]) {
				return false
			}
		}
		return true
	}
	isDescending := func(data []T) bool {
		for i := 0; i < size-1; i++ {
			if iteratee(data[i]) < iteratee(data[i+1]) {
				return false
			}
		}
		return true
	}
	return isAscending(slice) || isDescending(slice)
}
