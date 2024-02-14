package conv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, String(tt.args.any), "String(%v)", tt.args.any)
		})
	}
}

func TestUtf8ToGbk(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "",
			args:    args{},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Utf8ToGbk(tt.args.bs)
			if !tt.wantErr(t, err, fmt.Sprintf("Utf8ToGbk(%v)", tt.args.bs)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Utf8ToGbk(%v)", tt.args.bs)
		})
	}
}

func TestGbkToUtf8(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "",
			args:    args{},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GbkToUtf8(tt.args.bs)
			if !tt.wantErr(t, err, fmt.Sprintf("GbkToUtf8(%v)", tt.args.bs)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GbkToUtf8(%v)", tt.args.bs)
		})
	}
}
