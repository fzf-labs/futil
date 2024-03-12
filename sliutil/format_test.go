package sliutil

import (
	"reflect"
	"testing"
)

type SliDemo1 struct {
	name string
}
type SliDemo2 struct {
	Name string
	M    map[string]string
}

var s1 = []int{1, 2, 3, 4, 5}
var s2 = &SliDemo1{
	name: "jo",
}
var s3 = &SliDemo2{
	Name: "jo",
	M:    map[string]string{"a": "b"},
}

func TestNilSliceToEmptySlice(t *testing.T) {
	type args struct {
		inter any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "case 1",
			args: args{
				inter: s1,
			},
			want: s1,
		},
		{
			name: "case 2",
			args: args{
				inter: s2,
			},
			want: &SliDemo1{},
		},
		{
			name: "case 3",
			args: args{
				inter: s3,
			},
			want: s3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NilSliceToEmptySlice(tt.args.inter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilSliceToEmptySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
