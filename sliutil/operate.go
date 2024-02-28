package sliutil

import (
	"fmt"
	"math/rand"
	"strings"
)

// Contain 检查目标值是否在切片中。
func Contain[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// ContainBy 检查目标值是否满足函数。
func ContainBy[T any](slice []T, predicate func(item T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

// ContainSubSlice 检查切片是否包含给定的子切片。
func ContainSubSlice[T comparable](slice, subSlice []T) bool {
	for _, v := range subSlice {
		if !Contain(slice, v) {
			return false
		}
	}

	return true
}

// Chunk 创建一个元素切片，将其分成大小大小的组。
func Chunk[T any](slice []T, size int) [][]T {
	result := [][]T{}

	if len(slice) == 0 || size <= 0 {
		return result
	}

	for _, item := range slice {
		l := len(result)
		if l == 0 || len(result[l-1]) == size {
			result = append(result, []T{})
			l++
		}

		result[l-1] = append(result[l-1], item)
	}

	return result
}

// Compact 创建一个删除所有假值的切片。值false、nil、0和""为false。
func Compact[T comparable](slice []T) []T {
	var zero T
	var result []T
	for _, v := range slice {
		if v != zero {
			result = append(result, v)
		}
	}
	return result
}

// Concat 创建一个新的切片，将切片与任何其他切片连接起来。
func Concat[T any](slice []T, slices ...[]T) []T {
	result := append([]T{}, slice...)

	for _, v := range slices {
		result = append(result, v...)
	}

	return result
}

// Difference 获取差集
func Difference[T comparable](slice, comparedSlice []T) []T {
	var result []T

	for _, v := range slice {
		if !Contain(comparedSlice, v) {
			result = append(result, v)
		}
	}

	return result
}

// Equal 检查两个切片是否相等:长度相同，所有元素的顺序和值是否相等。
func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// Every 如果片中的所有值都通过回调函数，则返回true。
func Every[T any](slice []T, predicate func(index int, item T) bool) bool {
	for i, v := range slice {
		if !predicate(i, v) {
			return false
		}
	}
	return true
}

// None 如果片中的所有值都不符合标准，则返回true。
func None[T any](slice []T, predicate func(index int, item T) bool) bool {
	l := 0
	for i, v := range slice {
		if !predicate(i, v) {
			l++
		}
	}

	return l == len(slice)
}

// Some 如果列表中的任何值通过谓词函数，则返回true。
func Some[T any](slice []T, predicate func(index int, item T) bool) bool {
	for i, v := range slice {
		if predicate(i, v) {
			return true
		}
	}

	return false
}

// Filter 迭代slice的元素，返回传递谓词函数的所有元素的切片。
func Filter[T any](slice []T, predicate func(index int, item T) bool) []T {
	result := make([]T, 0)

	for i, v := range slice {
		if predicate(i, v) {
			result = append(result, v)
		}
	}

	return result
}

// FilterMap 返回对给定片应用过滤和映射的片。
func FilterMap[T any, U any](slice []T, iteratee func(index int, item T) (U, bool)) []U {
	var result []U
	for i, v := range slice {
		if a, ok := iteratee(i, v); ok {
			result = append(result, a)
		}
	}
	return result
}

// Count 返回给定项在片中出现的次数。
func Count[T comparable](slice []T, item T) int {
	count := 0

	for _, v := range slice {
		if item == v {
			count++
		}
	}

	return count
}

// CountBy 使用谓词函数遍历slice的元素，返回所有匹配元素的个数。
func CountBy[T any](slice []T, predicate func(index int, item T) bool) int {
	count := 0

	for i, v := range slice {
		if predicate(i, v) {
			count++
		}
	}

	return count
}

// GroupBy 迭代片的元素，每个元素将按标准分组，返回两个片。
func GroupBy[T any](slice []T, groupFn func(index int, item T) bool) ([]T, []T) {
	if len(slice) == 0 {
		return make([]T, 0), make([]T, 0)
	}

	groupB := make([]T, 0)
	groupA := make([]T, 0)

	for i, v := range slice {
		ok := groupFn(i, v)
		if ok {
			groupA = append(groupA, v)
		} else {
			groupB = append(groupB, v)
		}
	}

	return groupA, groupB
}

// GroupWith 返回由运行slice thru迭代器的每个元素的结果生成的键组成的映射。
func GroupWith[T any, U comparable](slice []T, iteratee func(item T) U) map[U][]T {
	result := make(map[U][]T)

	for _, v := range slice {
		key := iteratee(v)
		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		result[key] = append(result[key], v)
	}

	return result
}

// First 返回第一个元素，如果没有元素则返回零值。
func First[T any](ss []T) T {
	if len(ss) == 0 {
		var zeroValue T
		return zeroValue
	}
	return ss[0]
}

// Last 返回最后一个元素，如果没有元素则返回0值。
func Last[T any](ss []T) T {
	if len(ss) == 0 {
		var zeroValue T
		return zeroValue
	}
	return ss[len(ss)-1]
}

// Bottom 将从底部返回n个元素
func Bottom[T any](ss []T, n int) (top []T) {
	var lastIndex = len(ss) - 1
	for i := lastIndex; i > -1 && n > 0; i-- {
		top = append(top, ss[i])
		n--
	}

	return
}

// FindBy 迭代slice的元素，返回第一个通过谓词函数真值测试的元素。
func FindBy[T any](slice []T, predicate func(index int, item T) bool) (v T, ok bool) {
	index := -1

	for i, v := range slice {
		if predicate(i, v) {
			index = i
			break
		}
	}

	if index == -1 {
		return v, false
	}

	return slice[index], true
}

// FindLastBy 迭代slice的元素，返回最后一个通过谓词函数真值测试的元素。
func FindLastBy[T any](slice []T, predicate func(index int, item T) bool) (v T, ok bool) {
	index := -1
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(i, slice[i]) {
			index = i
			break
		}
	}
	if index == -1 {
		return v, false
	}
	return slice[index], true
}

// Map 通过运行slice thru迭代函数的每个元素来创建一个值片。
func Map[T any, U any](slice []T, iteratee func(index int, item T) U) []U {
	result := make([]U, len(slice), cap(slice))
	for i := 0; i < len(slice); i++ {
		result[i] = iteratee(i, slice[i])
	}
	return result
}

// Replace 返回切片的副本，其中将旧的前n个不重叠的实例替换为新实例。
func Replace[T comparable](slice []T, old T, new T, n int) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll 返回切片的副本，其中所有不重叠的old实例替换为new实例。
func ReplaceAll[T comparable](slice []T, old T, new T) []T {
	return Replace(slice, old, new, -1)
}

// Repeat 创建一个长度为n的切片，其元素参数为“item”。
// Play: https://go.dev/play/p/1CbOmtgILUU
func Repeat[T any](item T, n int) []T {
	result := make([]T, n)

	for i := range result {
		result[i] = item
	}

	return result
}

// DeleteAt 删除索引处切片的元素。
func DeleteAt[T any](slice []T, index int) []T {
	if index >= len(slice) {
		index = len(slice) - 1
	}

	result := make([]T, len(slice)-1)
	copy(result, slice[:index])
	copy(result[index:], slice[index+1:])

	return result
}

// DeleteRange 删除从开始索引到结束索引(排除)的slice元素。
func DeleteRange[T any](slice []T, start, end int) []T {
	result := make([]T, 0, len(slice)-(end-start))

	for i := 0; i < start; i++ {
		result = append(result, slice[i])
	}

	for i := end; i < len(slice); i++ {
		result = append(result, slice[i])
	}

	return result
}

// Drop 从切片的开始处删除n个元素。
func Drop[T any](slice []T, n int) []T {
	size := len(slice)

	if size <= n {
		return []T{}
	}

	if n <= 0 {
		return slice
	}

	result := make([]T, 0, size-n)

	return append(result, slice[n:]...)
}

// DropRight 从切片的末尾删除n个元素。
func DropRight[T any](slice []T, n int) []T {
	size := len(slice)
	if size <= n {
		return []T{}
	}
	if n <= 0 {
		return slice
	}
	result := make([]T, 0, size-n)
	return append(result, slice[:size-n]...)
}

// InsertAt 将值或其他切片插入到索引处的切片中。
func InsertAt[T any](slice []T, index int, value any) []T {
	size := len(slice)

	if index < 0 || index > size {
		return slice
	}

	if v, ok := value.(T); ok {
		slice = append(slice[:index], append([]T{v}, slice[index:]...)...)
		return slice
	}

	if v, ok := value.([]T); ok {
		slice = append(slice[:index], append(v, slice[index:]...)...)
		return slice
	}

	return slice
}

// UpdateAt 更新索引处的切片元素。
func UpdateAt[T any](slice []T, index int, value T) []T {
	size := len(slice)
	if index < 0 || index >= size {
		return slice
	}
	slice = append(slice[:index], append([]T{value}, slice[index+1:]...)...)
	return slice
}

// Unique 删除切片中的重复元素。
func Unique[T comparable](slice []T) []T {
	result := []T{}

	for i := 0; i < len(slice); i++ {
		v := slice[i]
		skip := true
		for j := range result {
			if v == result[j] {
				skip = false
				break
			}
		}
		if skip {
			result = append(result, v)
		}
	}

	return result
}

// UniqueBy 对slice的每一项调用iteratee函数，然后删除重复项。
func UniqueBy[T comparable](slice []T, iteratee func(item T) T) []T {
	var result []T
	for _, v := range slice {
		val := iteratee(v)
		result = append(result, val)
	}
	return Unique(result)
}

// Union 从所有给定的片中按顺序创建唯一元素的片。
func Union[T comparable](slices ...[]T) []T {
	var result []T
	contain := map[T]struct{}{}
	for _, slice := range slices {
		for _, item := range slice {
			if _, ok := contain[item]; !ok {
				contain[item] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

// UnionBy 类似于Union，更重要的是，它接受iteritere，每个slice的每个元素都会被调用。
func UnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T {
	result := []T{}
	contain := map[V]struct{}{}
	for _, slice := range slices {
		for _, item := range slice {
			val := predicate(item)
			if _, ok := contain[val]; !ok {
				contain[val] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

// Merge 所有给定的切片合为一片。
func Merge[T any](slices ...[]T) []T {
	result := make([]T, 0)

	for _, v := range slices {
		result = append(result, v...)
	}

	return result
}

// Intersection 创建包含所有片的唯一元素的片。
func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
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

	result := reducer(slices[0], slices[1])

	reduceSlice := make([][]T, 2)
	for i := 2; i < len(slices); i++ {
		reduceSlice[0] = result
		reduceSlice[1] = slices[i]
		result = reducer(reduceSlice[0], reduceSlice[1])
	}

	return result
}

// SymmetricDifference 交点函数的反运算。
func SymmetricDifference[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
	}

	result := make([]T, 0)

	intersectSlice := Intersection(slices...)

	for i := 0; i < len(slices); i++ {
		slice := slices[i]
		for _, v := range slice {
			if !Contain(intersectSlice, v) {
				result = append(result, v)
			}
		}

	}

	return Unique(result)
}

// Reverse 返回元素顺序与给定切片相反的切片。
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Without 创建一个不包括所有给定项的片。
func Without[T comparable](slice []T, items ...T) []T {
	if len(items) == 0 || len(slice) == 0 {
		return slice
	}

	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !Contain(items, v) {
			result = append(result, v)
		}
	}

	return result
}

// IndexOf 返回在片中找到第一个条目的索引，如果找不到该条目则返回-1。
func IndexOf[T comparable](arr []T, val T) int {
	limit := 10
	// gets the hash value of the array as the key of the hash table.
	key := fmt.Sprintf("%p", arr)
	memoryHashMap := make(map[string]map[any]int)
	memoryHashCounter := make(map[string]int)
	// determines whether the hash table is empty. If so, the hash table is created.
	if memoryHashMap[key] == nil {
		memoryHashMap[key] = make(map[any]int)
		// iterate through the array, adding the value and index of each element to the hash table.
		for i := len(arr) - 1; i >= 0; i-- {
			memoryHashMap[key][arr[i]] = i
		}
	}
	// update the hash table counter.
	memoryHashCounter[key]++

	// use the hash table to find the specified value. If found, the index is returned.
	if index, ok := memoryHashMap[key][val]; ok {
		// calculate the memory usage of the hash table.
		size := len(memoryHashMap)
		// If the memory usage of the hash table exceeds the memory limit, the hash table with the lowest counter is cleared.
		if size > limit {
			var minKey string
			var minVal int
			for k, v := range memoryHashCounter {
				if k == key {
					continue
				}
				if minVal == 0 || v < minVal {
					minKey = k
					minVal = v
				}
			}
			delete(memoryHashMap, minKey)
			delete(memoryHashCounter, minKey)
		}
		return index
	}
	return -1
}

// LastIndexOf 返回在片中找到该项最后出现的索引，如果找不到该索引则返回-1。
func LastIndexOf[T comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if item == slice[i] {
			return i
		}
	}

	return -1
}

// ToSlicePointer 返回指向变量参数转换的切片的指针。
func ToSlicePointer[T any](items ...T) []*T {
	result := make([]*T, len(items))
	for i := range items {
		result[i] = &items[i]
	}

	return result
}

// ToSlice 返回变量参数转换的切片。
// Play: https://go.dev/play/p/YzbzVq5kscN
func ToSlice[T any](items ...T) []T {
	result := make([]T, len(items))
	copy(result, items)
	return result
}

// AppendIfAbsent 只有缺席的项目追加。
func AppendIfAbsent[T comparable](slice []T, item T) []T {
	if !Contain(slice, item) {
		slice = append(slice, item)
	}
	return slice
}

// KeyBy 基于回调函数将切片转换为映射。
func KeyBy[T any, U comparable](slice []T, iteratee func(item T) U) map[U]T {
	result := make(map[U]T, len(slice))

	for _, v := range slice {
		k := iteratee(v)
		result[k] = v
	}

	return result
}

// Join 带有指定分隔符的切片项。
func Join[T any](slice []T, separator string) string {
	str := Map(slice, func(_ int, item T) string {
		return fmt.Sprint(item)
	})
	return strings.Join(str, separator)
}

// Partition 所有具有给定谓词函数求值的片元素。
func Partition[T any](slice []T, predicates ...func(item T) bool) [][]T {
	l := len(predicates)
	result := make([][]T, l+1)
	for _, item := range slice {
		processed := false

		for i, f := range predicates {
			if f == nil {
				panic("predicate function must not be nill")
			}

			if f(item) {
				result[i] = append(result[i], item)
				processed = true
				break
			}
		}

		if !processed {
			result[l] = append(result[l], item)
		}
	}

	return result
}

// Random 获取slice的随机项，当slice为空时返回idx=-1
func Random[T any](slice []T) (val T, idx int) {
	if len(slice) == 0 {
		return val, -1
	}
	idx = rand.Intn(len(slice))
	return slice[idx], idx
}
