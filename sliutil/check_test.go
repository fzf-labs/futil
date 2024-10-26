package sliutil

import "testing"

func TestContain(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		target     T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case true",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				target:     1,
			},
			want: true,
		},
		{
			name: "case false",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				target:     6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.collection, tt.args.target); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainBy(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(item T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case true",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(item int) bool { return item == 1 },
			},
			want: true,
		},
		{
			name: "case false",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(item int) bool { return item == 6 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainBy(tt.args.collection, tt.args.predicate); got != tt.want {
				t.Errorf("ContainBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainSubSlice(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		subSlice   []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case true",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				subSlice:   []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "case false",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				subSlice:   []int{1, 2, 3, 6},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainSubSlice(tt.args.collection, tt.args.subSlice); got != tt.want {
				t.Errorf("ContainSubSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args[T comparable] struct {
		collection1 []T
		collection2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case true",
			args: args[int]{
				collection1: []int{1, 2, 3, 4, 5},
				collection2: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "case false 1",
			args: args[int]{
				collection1: []int{1, 2, 3, 4, 5},
				collection2: []int{1, 2, 3, 4, 6},
			},
			want: false,
		},
		{
			name: "case false 2",
			args: args[int]{
				collection1: []int{1, 2, 3, 4, 5},
				collection2: []int{1, 2, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.collection1, tt.args.collection2); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvery(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(index int, item T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case true",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(_ int, item int) bool { return item > 0 },
			},
			want: true,
		},
		{
			name: "case false",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(_ int, item int) bool { return item > 2 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Every(tt.args.collection, tt.args.predicate); got != tt.want {
				t.Errorf("Every() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNone(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(index int, item T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case true",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate: func(_ int, item int) bool {
					return item > 5
				},
			},
			want: true,
		},
		{
			name: "case false",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate: func(_ int, item int) bool {
					return item > 4
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := None(tt.args.collection, tt.args.predicate); got != tt.want {
				t.Errorf("None() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(index int, item T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case true",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(_ int, item int) bool { return item > 4 },
			},
			want: true,
		},
		{
			name: "case false",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(_ int, item int) bool { return item > 6 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Some(tt.args.collection, tt.args.predicate); got != tt.want {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}
