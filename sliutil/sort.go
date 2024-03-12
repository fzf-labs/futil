package sliutil

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

// Shuffle the slice.
func Shuffle[T any](collection []T) []T {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
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
