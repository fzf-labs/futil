package sliutil

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args[T any] struct {
		collection []T
		size       int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				size:       2,
			},
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.collection, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompact(t *testing.T) {
	type args[T comparable] struct {
		collection []T
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
				collection: []int{1, 2, 3, 4, 5, 0},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compact(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	type args[T any] struct {
		collection []T
		others     [][]T
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
				collection: []int{1, 2, 3, 4, 5},
				others:     [][]int{{6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.collection, tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(index int, item T) bool
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
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(index int, item int) bool { return item > 3 },
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.collection, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	type args[T any, U any] struct {
		collection []T
		iteratee   func(index int, item T) (U, bool)
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[int, int]{
		{
			name: "case 1",
			args: args[int, int]{
				collection: []int{1, 2, 3, 4, 5},
				iteratee:   func(index int, item int) (int, bool) { return item * 2, true },
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterMap(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	type args[T any] struct {
		collection []T
		groupFn    func(index int, item T) bool
	}
	type testCase[T any] struct {
		name  string
		args  args[T]
		want  []T
		want1 []T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				groupFn:    func(index int, item int) bool { return item > 3 },
			},
			want:  []int{4, 5},
			want1: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GroupBy(tt.args.collection, tt.args.groupFn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GroupBy() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGroupWith(t *testing.T) {
	type args[T any, U comparable] struct {
		collection []T
		iteratee   func(item T) U
	}
	type testCase[T any, U comparable] struct {
		name string
		args args[T, U]
		want map[U][]T
	}
	tests := []testCase[int, int]{
		{
			name: "case 1",
			args: args[int, int]{
				collection: []int{1, 2, 3, 4, 5},
				iteratee:   func(item int) int { return item % 2 },
			},
			want: map[int][]int{
				0: {2, 4},
				1: {1, 3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupWith(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	type args[T any] struct {
		collection []T
	}
	type testCase[T any] struct {
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
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := First(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args[T any] struct {
		collection []T
	}
	type testCase[T any] struct {
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
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Last(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBottom(t *testing.T) {
	type args[T any] struct {
		collection []T
		n          int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantTop []T
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				n:          2,
			},
			wantTop: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTop := Bottom(tt.args.collection, tt.args.n); !reflect.DeepEqual(gotTop, tt.wantTop) {
				t.Errorf("Bottom() = %v, want %v", gotTop, tt.wantTop)
			}
		})
	}
}

func TestFindBy(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(index int, item T) bool
	}
	type testCase[T any] struct {
		name   string
		args   args[T]
		wantV  T
		wantOk bool
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(index int, item int) bool { return item > 5 },
			},
			wantV:  0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotOk := FindBy(tt.args.collection, tt.args.predicate)
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("FindBy() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FindBy() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestFindLastBy(t *testing.T) {
	type args[T any] struct {
		collection []T
		predicate  func(index int, item T) bool
	}
	type testCase[T any] struct {
		name   string
		args   args[T]
		wantV  T
		wantOk bool
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				predicate:  func(index int, item int) bool { return item > 5 },
			},
			wantV:  0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotOk := FindLastBy(tt.args.collection, tt.args.predicate)
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("FindLastBy() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FindLastBy() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		collection []T
		iteratee   func(index int, item T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[int, int]{
		{
			name: "case 1",
			args: args[int, int]{
				collection: []int{1, 2, 3, 4, 5},
				iteratee:   func(index int, item int) int { return item * 2 },
			},
			want: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		old        T
		new        T
		n          int
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
				collection: []int{1, 2, 3, 4, 5},
				old:        1,
				new:        6,
				n:          1,
			},
			want: []int{6, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.args.collection, tt.args.old, tt.args.new, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceAll(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		old        T
		new        T
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
				collection: []int{1, 2, 3, 4, 5},
				old:        1,
				new:        6,
			},
			want: []int{6, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAll(tt.args.collection, tt.args.old, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReplaceAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args[T any] struct {
		item T
		n    int
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
				item: 1,
				n:    5,
			},
			want: []int{1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.item, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimes(t *testing.T) {
	type args[T any] struct {
		count    int
		iteratee func(index int) T
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
				count:    5,
				iteratee: func(index int) int { return index },
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Times(tt.args.count, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	type args[T any] struct {
		collection []T
		index      int
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
				collection: []int{1, 2, 3, 4, 5},
				index:      1,
			},
			want: []int{1, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteAt(tt.args.collection, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRange(t *testing.T) {
	type args[T any] struct {
		collection []T
		start      int
		end        int
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
				collection: []int{1, 2, 3, 4, 5},
				start:      1,
				end:        3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteRange(tt.args.collection, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrop(t *testing.T) {
	type args[T any] struct {
		collection []T
		n          int
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
				collection: []int{1, 2, 3, 4, 5},
				n:          2,
			},
			want: []int{3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Drop(tt.args.collection, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Drop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropRight(t *testing.T) {
	type args[T any] struct {
		collection []T
		n          int
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
				collection: []int{1, 2, 3, 4, 5},
				n:          2,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropRight(tt.args.collection, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	type args[T any] struct {
		collection []T
		index      int
		value      any
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
				collection: []int{1, 2, 3, 4, 5},
				index:      2,
				value:      6,
			},
			want: []int{1, 2, 6, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAt(tt.args.collection, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateAt(t *testing.T) {
	type args[T any] struct {
		collection []T
		index      int
		value      T
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
				collection: []int{1, 2, 3, 4, 5},
				index:      2,
				value:      6,
			},
			want: []int{1, 2, 6, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateAt(tt.args.collection, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args[T comparable] struct {
		collection []T
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
				collection: []int{1, 2, 3, 4, 5},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueBy(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		iteratee   func(item T) T
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
				collection: []int{1, 2, 3, 4, 5},
				iteratee:   func(item int) int { return item },
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueBy(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args[T any] struct {
		collections [][]T
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
				collections: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
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
				collections: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionBy(t *testing.T) {
	type args[T any, V comparable] struct {
		predicate func(item T) V
		slices    [][]T
	}
	type testCase[T any, V comparable] struct {
		name string
		args args[T, V]
		want []T
	}
	tests := []testCase[int, int]{
		{
			name: "case 1",
			args: args[int, int]{
				predicate: func(item int) int { return item },
				slices:    [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionBy(tt.args.predicate, tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
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
				collection: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.args.collection)
			if !reflect.DeepEqual(tt.args.collection, tt.want) {
				t.Errorf("Reverse() = %v, want %v", tt.args.collection, tt.want)
			}
		})
	}
}

func TestWithout(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		items      []T
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
				collection: []int{1, 2, 3, 4, 5},
				items:      []int{1, 2},
			},
			want: []int{3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Without(tt.args.collection, tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Without() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		val        T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.collection, tt.args.val); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndexOf(t *testing.T) {
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
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIndexOf(tt.args.collection, tt.args.item); got != tt.want {
				t.Errorf("LastIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlicePointer(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	type args[T any] struct {
		collections []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []*T
	}
	tests := []testCase[User]{
		{
			name: "case 1",
			args: args[User]{
				collections: []User{{Name: "Tom", Age: 18}, {Name: "Jerry", Age: 20}},
			},
			want: []*User{{Name: "Tom", Age: 18}, {Name: "Jerry", Age: 20}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSlicePointer(tt.args.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlicePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	type args[T any] struct {
		collections []T
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
				collections: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSlice(tt.args.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendIfAbsent(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		item       T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "case	1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				item:       0,
			},
			want: []int{1, 2, 3, 4, 5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendIfAbsent(tt.args.collection, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendIfAbsent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyBy(t *testing.T) {
	type args[T any, U comparable] struct {
		collection []T
		iteratee   func(item T) U
	}
	type testCase[T any, U comparable] struct {
		name string
		args args[T, U]
		want map[U]T
	}
	tests := []testCase[int, int]{
		{
			name: "case 1",
			args: args[int, int]{
				collection: []int{1, 2, 3, 4, 5},
				iteratee:   func(item int) int { return item },
			},
			want: map[int]int{
				1: 1,
				2: 2,
				3: 3,
				4: 4,
				5: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyBy(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	type args[T any] struct {
		collection []T
		separator  string
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
				separator:  ",",
			},
			want: "1,2,3,4,5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.collection, tt.args.separator); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	type args[T any] struct {
		collection []T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantVal T
		wantIdx int
	}
	tests := []testCase[int]{
		{
			name: "case 1",
			args: args[int]{
				collection: []int{1, 2, 3, 4, 5},
			},
			wantVal: 1,
			wantIdx: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotIdx := Random(tt.args.collection)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Random() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotIdx != tt.wantIdx {
				t.Errorf("Random() gotIdx = %v, want %v", gotIdx, tt.wantIdx)
			}
		})
	}
}
