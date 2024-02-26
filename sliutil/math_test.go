package sliutil

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestSum(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		ss []T
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
			name: "",
			args: args[int]{
				ss: []int{1, 2, 3},
			},
			wantSum: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.args.ss); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		ss []T
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
				ss: []int{1, 2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.ss); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args[T constraints.Ordered] struct {
		ss []T
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
				ss: []int{1, 2, 3},
			},
			wantMin: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := Max(tt.args.ss); !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("Max() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		ss []T
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
				ss: []int{1, 2, 3},
			},
			wantMin: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := Min(tt.args.ss); !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("Min() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		ss []T
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
				ss: []int{1, 2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.ss); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		ss []T
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
				ss: []int{1, 2, 3},
			},
			wantProduct: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProduct := Product(tt.args.ss); gotProduct != tt.wantProduct {
				t.Errorf("Product() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		ss []T
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
				ss: []int{1, 2, 3},
			},
			want: 0.816496580927726,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDeviation(tt.args.ss); got != tt.want {
				t.Errorf("StandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
