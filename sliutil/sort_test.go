package sliutil

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{
			name: "empty slice",
			arr:  []int{},
		},
		{
			name: "single element",
			arr:  []int{1},
		},
		{
			name: "multiple elements",
			arr:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.arr))
			copy(original, tt.arr)

			result := Shuffle(tt.arr)

			// 检查长度是否相同
			if len(result) != len(original) {
				t.Errorf("Shuffle() returned slice with different length, got %v, want %v", len(result), len(original))
			}

			// 检查元素是否相同（可能顺序不同）
			elementCount := make(map[int]int)
			for _, v := range original {
				elementCount[v]++
			}
			for _, v := range result {
				elementCount[v]--
				if elementCount[v] == 0 {
					delete(elementCount, v)
				}
			}
			if len(elementCount) != 0 {
				t.Errorf("Shuffle() returned slice with different elements")
			}
		})
	}
}

func TestIsAsc(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "empty slice",
			arr:  []int{},
			want: true,
		},
		{
			name: "single element",
			arr:  []int{1},
			want: true,
		},
		{
			name: "ascending order",
			arr:  []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "not ascending order",
			arr:  []int{1, 3, 2, 4, 5},
			want: false,
		},
		{
			name: "descending order",
			arr:  []int{5, 4, 3, 2, 1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAsc(tt.arr); got != tt.want {
				t.Errorf("IsAsc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDesc(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "empty slice",
			arr:  []int{},
			want: true,
		},
		{
			name: "single element",
			arr:  []int{1},
			want: true,
		},
		{
			name: "descending order",
			arr:  []int{5, 4, 3, 2, 1},
			want: true,
		},
		{
			name: "not descending order",
			arr:  []int{5, 3, 4, 2, 1},
			want: false,
		},
		{
			name: "ascending order",
			arr:  []int{1, 2, 3, 4, 5},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDesc(tt.arr); got != tt.want {
				t.Errorf("IsDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "empty slice",
			arr:  []int{},
			want: true,
		},
		{
			name: "single element",
			arr:  []int{1},
			want: true,
		},
		{
			name: "ascending order",
			arr:  []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "descending order",
			arr:  []int{5, 4, 3, 2, 1},
			want: true,
		},
		{
			name: "unsorted",
			arr:  []int{1, 3, 2, 5, 4},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSorted(tt.arr); got != tt.want {
				t.Errorf("IsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSortedByKey(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	tests := []struct {
		name     string
		arr      []person
		iteratee func(person) int
		want     bool
	}{
		{
			name: "empty slice",
			arr:  []person{},
			iteratee: func(p person) int {
				return p.age
			},
			want: true,
		},
		{
			name: "single element",
			arr:  []person{{name: "Alice", age: 20}},
			iteratee: func(p person) int {
				return p.age
			},
			want: true,
		},
		{
			name: "ascending order by age",
			arr: []person{
				{name: "Alice", age: 20},
				{name: "Bob", age: 25},
				{name: "Charlie", age: 30},
			},
			iteratee: func(p person) int {
				return p.age
			},
			want: true,
		},
		{
			name: "descending order by age",
			arr: []person{
				{name: "Charlie", age: 30},
				{name: "Bob", age: 25},
				{name: "Alice", age: 20},
			},
			iteratee: func(p person) int {
				return p.age
			},
			want: true,
		},
		{
			name: "unsorted by age",
			arr: []person{
				{name: "Alice", age: 20},
				{name: "Charlie", age: 30},
				{name: "Bob", age: 25},
			},
			iteratee: func(p person) int {
				return p.age
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSortedByKey(tt.arr, tt.iteratee); got != tt.want {
				t.Errorf("IsSortedByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
