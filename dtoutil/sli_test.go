package dto

import (
	"reflect"
	"testing"
)

type TestInterface interface {
	DoSomething()
}

type TestImplementation struct {
	Data []int
}

func (t *TestImplementation) DoSomething() {}

func TestNilSliceToEmptySlice(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    any
		wantErr bool
	}{
		{
			name:    "Nil input",
			input:   nil,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Nil slice",
			input:   []int(nil),
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "Non-nil slice",
			input:   []int{1, 2, 3},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "Nil map",
			input:   map[string]int(nil),
			want:    map[string]int{},
			wantErr: false,
		},
		{
			name:    "Non-nil map",
			input:   map[string]int{"a": 1},
			want:    map[string]int{"a": 1},
			wantErr: false,
		},
		{
			name:    "Nil pointer",
			input:   (*int)(nil),
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Non-nil pointer",
			input:   func() *int { i := 5; return &i }(),
			want:    func() *int { i := 5; return &i }(),
			wantErr: false,
		},
		{
			name:    "Nil interface",
			input:   interface{}(nil),
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Non-nil interface",
			input:   interface{}(5),
			want:    5,
			wantErr: false,
		},
		{
			name:    "Struct with nil slice",
			input:   struct{ S []int }{S: nil},
			want:    struct{ S []int }{S: []int{}},
			wantErr: false,
		},
		{
			name:    "Array",
			input:   [3]int{1, 2, 3},
			want:    [3]int{1, 2, 3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NilSliceToEmptySlice(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NilSliceToEmptySlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilSliceToEmptySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 测试递归深度限制
func TestNilSliceToEmptySlice_RecursionLimit(t *testing.T) {
	// 创建一个深度超过限制的结构
	type Node struct {
		Next *Node
		Data []int
	}

	root := &Node{Data: nil}
	current := root
	for i := 0; i < maxDepth+1; i++ {
		current.Next = &Node{Data: nil}
		current = current.Next
	}

	_, err := NilSliceToEmptySlice(root)
	if err == nil {
		t.Error("Expected error for exceeding max recursion depth")
	}
}

// 测试类型兼容性
func TestCanAssign(t *testing.T) {
	tests := []struct {
		name string
		from reflect.Type
		to   reflect.Type
		want bool
	}{
		{
			name: "same types",
			from: reflect.TypeOf(0),
			to:   reflect.TypeOf(0),
			want: true,
		},
		{
			name: "convertible types",
			from: reflect.TypeOf(int32(0)),
			to:   reflect.TypeOf(int64(0)),
			want: true,
		},
		{
			name: "pointer to non-pointer",
			from: reflect.TypeOf(&[]int{}),
			to:   reflect.TypeOf([]int{}),
			want: true,
		},
		{
			name: "interface implementation",
			from: reflect.TypeOf(&TestImplementation{}),
			to:   reflect.TypeOf((*TestInterface)(nil)).Elem(),
			want: true,
		},
		{
			name: "incompatible types",
			from: reflect.TypeOf("string"),
			to:   reflect.TypeOf(0),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAssign(tt.from, tt.to); got != tt.want {
				t.Errorf("canAssign() = %v, want %v", got, tt.want)
			}
		})
	}
}
