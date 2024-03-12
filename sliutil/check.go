package sliutil

// Contain 检查目标值是否在切片中。
func Contain[T comparable](collection []T, target T) bool {
	for _, item := range collection {
		if item == target {
			return true
		}
	}
	return false
}

// ContainBy 检查目标值是否满足函数。
func ContainBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// ContainSubSlice 检查切片是否包含给定的子切片。
func ContainSubSlice[T comparable](collection, subSlice []T) bool {
	for _, v := range subSlice {
		if !Contain(collection, v) {
			return false
		}
	}

	return true
}

// Equal 检查两个切片是否相等:长度相同，所有元素的顺序和值是否相等。
func Equal[T comparable](collection1, collection2 []T) bool {
	if len(collection1) != len(collection2) {
		return false
	}

	for i := range collection1 {
		if collection1[i] != collection2[i] {
			return false
		}
	}

	return true
}

// Every 如果片中的所有值都通过回调函数，则返回true。
func Every[T any](collection []T, predicate func(index int, item T) bool) bool {
	for i, v := range collection {
		if !predicate(i, v) {
			return false
		}
	}
	return true
}

// None 如果片中的所有值都不符合标准，则返回true。
func None[T any](collection []T, predicate func(index int, item T) bool) bool {
	l := 0
	for i, v := range collection {
		if !predicate(i, v) {
			l++
		}
	}

	return l == len(collection)
}

// Some 如果列表中的任何值通过谓词函数，则返回true。
func Some[T any](collection []T, predicate func(index int, item T) bool) bool {
	for i, v := range collection {
		if predicate(i, v) {
			return true
		}
	}
	return false
}
