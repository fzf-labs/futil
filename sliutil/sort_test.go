package sliutil

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestIsAsc(t *testing.T) {
	type args[T constraints.Ordered] struct {
		collection []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAsc(tt.args.collection); got != tt.want {
				t.Errorf("IsAsc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDesc(t *testing.T) {
	type args[T constraints.Ordered] struct {
		collection []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDesc(tt.args.collection); got != tt.want {
				t.Errorf("IsDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSorted(t *testing.T) {
	type args[T constraints.Ordered] struct {
		collection []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSorted(tt.args.collection); got != tt.want {
				t.Errorf("IsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSortedByKey(t *testing.T) {
	type args[T any, K constraints.Ordered] struct {
		collection []T
		iteratee   func(item T) K
	}
	type testCase[T any, K constraints.Ordered] struct {
		name string
		args args[T, K]
		want bool
	}
	tests := []testCase[int, int]{
		{
			name: "case 1",
			args: args[int, int]{
				collection: nil,
				iteratee:   nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSortedByKey(tt.args.collection, tt.args.iteratee); got != tt.want {
				t.Errorf("IsSortedByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	type args[T any] struct {
		collection []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shuffle(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
