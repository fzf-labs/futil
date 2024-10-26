package sliutil

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Sum 求和是所有元素的和。
func Sum[T constraints.Integer | constraints.Float](collection []T) (sum T) {
	for _, s := range collection {
		sum += s
	}
	return
}

// Average 元素的平均值，如果没有则为零
func Average[T constraints.Integer | constraints.Float](collection []T) float64 {
	if l := len(collection); l > 0 {
		return float64(Sum(collection)) / float64(l)
	}
	return 0
}

// Max 是最大值，或零。
func Max[T constraints.Ordered](collection []T) T {
	var maxItem T
	if len(collection) == 0 {
		return maxItem
	}
	maxItem = collection[0]
	for _, s := range collection {
		if s > maxItem {
			maxItem = s
		}
	}
	return maxItem
}

// Min 是最小值，或零。
func Min[T constraints.Ordered](collection []T) T {
	var minItem T
	if len(collection) == 0 {
		return minItem
	}
	minItem = collection[0]
	for _, s := range collection {
		if s < minItem {
			minItem = s
		}
	}
	return minItem
}

// Count 返回给定项在片中出现的次数。
func Count[T comparable](collection []T, item T) int {
	count := 0
	for _, v := range collection {
		if item == v {
			count++
		}
	}
	return count
}

// CountBy 使用谓词函数遍历slice的元素，返回所有匹配元素的个数。
func CountBy[T any](collection []T, predicate func(index int, item T) bool) int {
	count := 0
	for i, v := range collection {
		if predicate(i, v) {
			count++
		}
	}
	return count
}

// Intersection 交集
func Intersection[T comparable](collections ...[]T) []T {
	if len(collections) == 0 {
		return []T{}
	}
	if len(collections) == 1 {
		return Unique(collections[0])
	}
	reducer := func(sliceA, sliceB []T) []T {
		hashMap := make(map[T]int)
		for _, v := range sliceA {
			hashMap[v] = 1
		}

		out := make([]T, 0)
		for _, val := range sliceB {
			if v, ok := hashMap[val]; v == 1 && ok {
				out = append(out, val)
				hashMap[val]++
			}
		}
		return out
	}
	result := reducer(collections[0], collections[1])
	reduceSlice := make([][]T, 2)
	for i := 2; i < len(collections); i++ {
		reduceSlice[0] = result
		reduceSlice[1] = collections[i]
		result = reducer(reduceSlice[0], reduceSlice[1])
	}
	return result
}

// Difference 差集
func Difference[T comparable](collection, comparedSlice []T) []T {
	var result []T
	for _, v := range collection {
		if !Contain(comparedSlice, v) {
			result = append(result, v)
		}
	}
	return result
}

// SymmetricDifference 对称差
func SymmetricDifference[T comparable](collections ...[]T) []T {
	if len(collections) == 0 {
		return []T{}
	}
	if len(collections) == 1 {
		return Unique(collections[0])
	}
	result := make([]T, 0)
	intersectSlice := Intersection(collections...)
	for i := 0; i < len(collections); i++ {
		slice := collections[i]
		for _, v := range slice {
			if !Contain(intersectSlice, v) {
				result = append(result, v)
			}
		}
	}
	return Unique(result)
}

// Median  中位数
// 中位数返回数据样本的上半部分与下半部分之间的分隔值。如果切片中没有元素，则返回0。如果元素数量为偶数，则返回两个“中值”的ElementType平均值。
func Median[T constraints.Integer | constraints.Float](collection []T) T {
	n := len(collection)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return collection[0]
	}
	// This implementation aims at linear time O(n) on average.
	// It uses the same idea as QuickSort, but makes only 1 recursive
	// call instead of 2. See also Quickselect.
	work := make([]T, len(collection))
	copy(work, collection)
	limit1, limit2 := n/2, n/2+1
	if n%2 == 0 {
		limit1, limit2 = n/2-1, n/2+1
	}
	var rec func(a, b int)
	rec = func(a, b int) {
		if b-a <= 1 {
			return
		}
		ipivot := (a + b) / 2
		pivot := work[ipivot]
		work[a], work[ipivot] = work[ipivot], work[a]
		j := a
		k := b
		for j+1 < k {
			if work[j+1] < pivot {
				work[j+1], work[j] = work[j], work[j+1]
				j++
			} else {
				work[j+1], work[k-1] = work[k-1], work[j+1]
				k--
			}
		}
		// 1 or 0 recursive calls
		if j > limit1 {
			rec(a, j)
		}
		if j+1 < limit2 {
			rec(j+1, b)
		}
	}
	rec(0, len(work))
	if n%2 == 1 {
		return work[n/2]
	} else {
		return (work[n/2-1] + work[n/2]) / 2
	}
}

// Product 所有元素的乘积。
func Product[T constraints.Integer | constraints.Float](collection []T) (product T) {
	if len(collection) == 0 {
		return
	}
	product = collection[0]
	for _, s := range collection[1:] {
		product *= s
	}
	return
}

// StandardDeviation 标准差
func StandardDeviation[T constraints.Integer | constraints.Float](collection []T) float64 {
	if len(collection) == 0 {
		return 0.0
	}
	avg := Average(collection)
	var sd float64
	for i := range collection {
		sd += math.Pow(float64(collection[i])-avg, 2)
	}
	sd = math.Sqrt(sd / float64(len(collection)))
	return sd
}
