package sliutil

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestAverage(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		collection []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want float64
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.collection); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		item       T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				item:       1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.collection, tt.args.item); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountBy(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(index int, item T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate: func(index int, item int) bool {
					return item > 3
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBy(tt.args.collection, tt.args.predicate); got != tt.want {
				t.Errorf("CountBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args[T comparable] struct {
		collection    []T
		comparedSlice []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection:    []int{1, 2, 3, 4, 5},
				comparedSlice: []int{1, 2, 3},
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.collection, tt.args.comparedSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	type args[T comparable] struct {
		collections [][]T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collections: [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args[T constraints.Ordered] struct {
		collection []T
	}
	type testCase[T constraints.Ordered] struct {
		name    string
		args    args[T]
		wantMin T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			wantMin: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := Max(tt.args.collection); !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("Max() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		collection []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.collection); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		collection []T
	}
	type testCase[T constraints.Ordered] struct {
		name    string
		args    args[T]
		wantMin T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			wantMin: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := Min(tt.args.collection); !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("Min() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		collection []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name        string
		args        args[T]
		wantProduct T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			wantProduct: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProduct := Product(tt.args.collection); gotProduct != tt.wantProduct {
				t.Errorf("Product() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		collection []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want float64
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			want: 1.4142135623730951,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDeviation(tt.args.collection); got != tt.want {
				t.Errorf("StandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		collection []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name    string
		args    args[T]
		wantSum T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			wantSum: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.args.collection); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	type args[T comparable] struct {
		collections [][]T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collections: [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SymmetricDifference(tt.args.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SymmetricDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
