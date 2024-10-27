package sliutil

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// Chunk 创建一个元素切片，将其分成大小大小的组。
func Chunk[T any](collection []T, size int) [][]T {
	var result [][]T
	if len(collection) == 0 || size <= 0 {
		return result
	}
	for _, item := range collection {
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
func Compact[T comparable](collection []T) []T {
	var zero T
	var result []T
	for _, v := range collection {
		if v != zero {
			result = append(result, v)
		}
	}
	return result
}

// Concat 创建一个新的切片，将切片与任何其他切片连接起来。
func Concat[T any](collection []T, others ...[]T) []T {
	result := append([]T{}, collection...)
	for _, v := range others {
		result = append(result, v...)
	}
	return result
}

// Filter 迭代slice的元素，返回传递谓词函数的所有元素的切片。
func Filter[T any](collection []T, predicate func(index int, item T) bool) []T {
	result := make([]T, 0)

	for i, v := range collection {
		if predicate(i, v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterMap 返回对给定片应用过滤和映射的片。
func FilterMap[T any, U any](collection []T, iteratee func(index int, item T) (U, bool)) []U {
	var result []U
	for i, v := range collection {
		if a, ok := iteratee(i, v); ok {
			result = append(result, a)
		}
	}
	return result
}

// GroupBy 迭代片的元素，每个元素将按标准分组，返回两个片。
func GroupBy[T any](collection []T, groupFn func(index int, item T) bool) (groupA, groupB []T) {
	if len(collection) == 0 {
		return
	}
	for i, v := range collection {
		ok := groupFn(i, v)
		if ok {
			groupA = append(groupA, v)
		} else {
			groupB = append(groupB, v)
		}
	}
	return
}

// GroupWith 返回由运行slice thru迭代器的每个元素的结果生成的键组成的映射。
func GroupWith[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	result := make(map[U][]T)

	for _, v := range collection {
		key := iteratee(v)
		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		result[key] = append(result[key], v)
	}

	return result
}

// First 返回第一个元素，如果没有元素则返回零值。
func First[T any](collection []T) T {
	if len(collection) == 0 {
		var zeroValue T
		return zeroValue
	}
	return collection[0]
}

// Last 返回最后一个元素，如果没有元素则返回0值。
func Last[T any](collection []T) T {
	if len(collection) == 0 {
		var zeroValue T
		return zeroValue
	}
	return collection[len(collection)-1]
}

// Bottom 将从底部返回n个元素
func Bottom[T any](collection []T, n int) (top []T) {
	var lastIndex = len(collection) - 1
	for i := lastIndex; i > -1 && n > 0; i-- {
		top = append(top, collection[i])
		n--
	}

	return
}

// FindBy 迭代slice的元素，返回第一个通过谓词函数真值测试的元素。
func FindBy[T any](collection []T, predicate func(index int, item T) bool) (v T, ok bool) {
	index := -1

	for i, v := range collection {
		if predicate(i, v) {
			index = i
			break
		}
	}

	if index == -1 {
		return v, false
	}

	return collection[index], true
}

// FindLastBy 迭代slice的元素，返回最后一个通过谓词函数真值测试的元素。
func FindLastBy[T any](collection []T, predicate func(index int, item T) bool) (v T, ok bool) {
	index := -1
	for i := len(collection) - 1; i >= 0; i-- {
		if predicate(i, collection[i]) {
			index = i
			break
		}
	}
	if index == -1 {
		return v, false
	}
	return collection[index], true
}

// Map 通过运行slice thru迭代函数的每个元素来创建一个值片。
func Map[T any, U any](collection []T, iteratee func(index int, item T) U) []U {
	result := make([]U, len(collection), cap(collection))
	for i := 0; i < len(collection); i++ {
		result[i] = iteratee(i, collection[i])
	}
	return result
}

// Replace 返回切片的副本，其中将旧的前n个不重叠的实例替换为新实例。
func Replace[T comparable](collection []T, oldItem, newItem T, n int) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	for i := range result {
		if result[i] == oldItem && n != 0 {
			result[i] = newItem
			n--
		}
	}

	return result
}

// ReplaceAll 返回切片的副本，其中所有不重叠的old实例替换为new实例。
func ReplaceAll[T comparable](collection []T, oldItem, newItem T) []T {
	return Replace(collection, oldItem, newItem, -1)
}

// Repeat 创建一个长度为n的切片，其元素参数为“item”。
func Repeat[T any](item T, n int) []T {
	result := make([]T, n)

	for i := range result {
		result[i] = item
	}

	return result
}

// Times 调用迭代对象n次，返回每次调用结果的数组。
func Times[T any](count int, iteratee func(index int) T) []T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = iteratee(i)
	}
	return result
}

// DeleteAt 删除索引处切片的元素。
func DeleteAt[T any](collection []T, index int) []T {
	if index >= len(collection) {
		index = len(collection) - 1
	}

	result := make([]T, len(collection)-1)
	copy(result, collection[:index])
	copy(result[index:], collection[index+1:])

	return result
}

// DeleteRange 删除从开始索引到结束索引(排除)的slice元素。
func DeleteRange[T any](collection []T, start, end int) []T {
	result := make([]T, 0, len(collection)-(end-start))

	for i := 0; i < start; i++ {
		result = append(result, collection[i])
	}

	for i := end; i < len(collection); i++ {
		result = append(result, collection[i])
	}

	return result
}

// Drop 从切片的开始处删除n个元素。
func Drop[T any](collection []T, n int) []T {
	size := len(collection)

	if size <= n {
		return []T{}
	}

	if n <= 0 {
		return collection
	}

	result := make([]T, 0, size-n)

	return append(result, collection[n:]...)
}

// DropRight 从切片的末尾删除n个元素。
func DropRight[T any](collection []T, n int) []T {
	size := len(collection)
	if size <= n {
		return []T{}
	}
	if n <= 0 {
		return collection
	}
	result := make([]T, 0, size-n)
	return append(result, collection[:size-n]...)
}

// InsertAt 将值或其他切片插入到索引处的切片中。
func InsertAt[T any](collection []T, index int, value any) []T {
	size := len(collection)

	if index < 0 || index > size {
		return collection
	}

	if v, ok := value.(T); ok {
		collection = append(collection[:index], append([]T{v}, collection[index:]...)...)
		return collection
	}

	if v, ok := value.([]T); ok {
		collection = append(collection[:index], append(v, collection[index:]...)...)
		return collection
	}

	return collection
}

// UpdateAt 更新索引处的切片元素。
func UpdateAt[T any](collection []T, index int, value T) []T {
	size := len(collection)
	if index < 0 || index >= size {
		return collection
	}
	collection = append(collection[:index], append([]T{value}, collection[index+1:]...)...)
	return collection
}

// Unique 删除切片中的重复元素。
func Unique[T comparable](collection []T) []T {
	var result []T
	for i := 0; i < len(collection); i++ {
		v := collection[i]
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
func UniqueBy[T comparable](collection []T, iteratee func(item T) T) []T {
	var result []T
	for _, v := range collection {
		val := iteratee(v)
		result = append(result, val)
	}
	return Unique(result)
}

// Merge 合并
func Merge[T any](collections ...[]T) []T {
	result := make([]T, 0)

	for _, v := range collections {
		result = append(result, v...)
	}

	return result
}

// Union 去重
func Union[T comparable](collections ...[]T) []T {
	var result []T
	contain := map[T]struct{}{}
	for _, slice := range collections {
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
	var result []T
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

// Reverse 返回元素顺序与给定切片相反的切片。
func Reverse[T any](collection []T) {
	for i, j := 0, len(collection)-1; i < j; i, j = i+1, j-1 {
		collection[i], collection[j] = collection[j], collection[i]
	}
}

// Without 创建一个不包括所有给定项的片。
func Without[T comparable](collection []T, items ...T) []T {
	if len(items) == 0 || len(collection) == 0 {
		return collection
	}

	result := make([]T, 0, len(collection))
	for _, v := range collection {
		if !Contain(items, v) {
			result = append(result, v)
		}
	}

	return result
}

// IndexOf 返回在片中找到第一个条目的索引，如果找不到该条目则返回-1。
func IndexOf[T comparable](collection []T, val T) int {
	limit := 10
	// gets the hash value of the array as the key of the hash table.
	key := fmt.Sprintf("%p", collection)
	memoryHashMap := make(map[string]map[any]int)
	memoryHashCounter := make(map[string]int)
	// determines whether the hash table is empty. If so, the hash table is created.
	if memoryHashMap[key] == nil {
		memoryHashMap[key] = make(map[any]int)
		// iterate through the array, adding the value and index of each element to the hash table.
		for i := len(collection) - 1; i >= 0; i-- {
			memoryHashMap[key][collection[i]] = i
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
func LastIndexOf[T comparable](collection []T, item T) int {
	for i := len(collection) - 1; i >= 0; i-- {
		if item == collection[i] {
			return i
		}
	}

	return -1
}

// ToSlicePointer 返回指向变量参数转换的切片的指针。
func ToSlicePointer[T any](collections ...T) []*T {
	result := make([]*T, len(collections))
	for i := range collections {
		result[i] = &collections[i]
	}

	return result
}

// ToSlice 返回变量参数转换的切片。
func ToSlice[T any](collections ...T) []T {
	result := make([]T, len(collections))
	copy(result, collections)
	return result
}

// AppendIfAbsent 只有缺席的项目追加。
func AppendIfAbsent[T comparable](collection []T, item T) []T {
	if !Contain(collection, item) {
		collection = append(collection, item)
	}
	return collection
}

// KeyBy 基于回调函数将切片转换为映射。
func KeyBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U]T {
	result := make(map[U]T, len(collection))

	for _, v := range collection {
		k := iteratee(v)
		result[k] = v
	}

	return result
}

// Join 带有指定分隔符的切片项。
func Join[T any](collection []T, separator string) string {
	str := Map(collection, func(_ int, item T) string {
		return fmt.Sprint(item)
	})
	return strings.Join(str, separator)
}

// Random 获取slice的随机项，当slice为空时返回idx=-1
func Random[T any](collection []T) (val T, idx int) {
	if len(collection) == 0 {
		return val, -1
	}
	// Using non-cryptographic random for general utility randomization
	idx = rand.IntN(len(collection))
	return collection[idx], idx
}
