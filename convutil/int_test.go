package conv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				any: nil,
			},
			want:    0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Int(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Int(%v)", tt.args.any)
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "",
			args:    args{},
			want:    0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Int8(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Int8(%v)", tt.args.any)
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "",
			args:    args{},
			want:    0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Int16(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Int16(%v)", tt.args.any)
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "",
			args:    args{},
			want:    0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Int32(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Int32(%v)", tt.args.any)
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				any: 'a',
			},
			want:    97,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Int64(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Int64(%v)", tt.args.any)
		})
	}
}
