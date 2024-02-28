package sliutil

import (
	"reflect"
	"testing"
)

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
			name: "",
			args: args{},
			want: nil,
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
