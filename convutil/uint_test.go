package conv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case nil",
			args: args{
				any: nil,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case true",
			args: args{
				any: true,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "case false",
			args: args{
				any: false,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case int",
			args: args{
				any: int(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int8",
			args: args{
				any: int8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int16",
			args: args{
				any: int16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int32",
			args: args{
				any: int32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int64",
			args: args{
				any: int64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint",
			args: args{
				any: uint(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint8",
			args: args{
				any: uint8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint16",
			args: args{
				any: uint16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint32",
			args: args{
				any: uint32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint64",
			args: args{
				any: uint64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float32",
			args: args{
				any: float32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float64",
			args: args{
				any: float64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case string",
			args: args{
				any: "8",
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case []byte",
			args: args{
				any: []byte("8"),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case struct",
			args: args{
				any: struct{}{},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Uint(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Uint(%v)", tt.args.any)
		})
	}
}

func TestUint8(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case nil",
			args: args{
				any: nil,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case true",
			args: args{
				any: true,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "case false",
			args: args{
				any: false,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case int",
			args: args{
				any: int(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int8",
			args: args{
				any: int8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int16",
			args: args{
				any: int16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int32",
			args: args{
				any: int32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int64",
			args: args{
				any: int64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint",
			args: args{
				any: uint(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint8",
			args: args{
				any: uint8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint16",
			args: args{
				any: uint16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint32",
			args: args{
				any: uint32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint64",
			args: args{
				any: uint64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float32",
			args: args{
				any: float32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float64",
			args: args{
				any: float64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case string",
			args: args{
				any: "8",
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case []byte",
			args: args{
				any: []byte("8"),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case struct",
			args: args{
				any: struct{}{},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint8(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Uint8(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Uint8(%v)", tt.args.any)
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case nil",
			args: args{
				any: nil,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case true",
			args: args{
				any: true,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "case false",
			args: args{
				any: false,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case int",
			args: args{
				any: int(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int8",
			args: args{
				any: int8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int16",
			args: args{
				any: int16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int32",
			args: args{
				any: int32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int64",
			args: args{
				any: int64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint",
			args: args{
				any: uint(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint8",
			args: args{
				any: uint8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint16",
			args: args{
				any: uint16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint32",
			args: args{
				any: uint32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint64",
			args: args{
				any: uint64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float32",
			args: args{
				any: float32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float64",
			args: args{
				any: float64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case string",
			args: args{
				any: "8",
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case []byte",
			args: args{
				any: []byte("8"),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case struct",
			args: args{
				any: struct{}{},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint16(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Uint16(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Uint16(%v)", tt.args.any)
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case nil",
			args: args{
				any: nil,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case true",
			args: args{
				any: true,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "case false",
			args: args{
				any: false,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case int",
			args: args{
				any: int(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int8",
			args: args{
				any: int8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int16",
			args: args{
				any: int16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int32",
			args: args{
				any: int32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int64",
			args: args{
				any: int64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint",
			args: args{
				any: uint(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint8",
			args: args{
				any: uint8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint16",
			args: args{
				any: uint16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint32",
			args: args{
				any: uint32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint64",
			args: args{
				any: uint64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float32",
			args: args{
				any: float32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float64",
			args: args{
				any: float64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case string",
			args: args{
				any: "8",
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case []byte",
			args: args{
				any: []byte("8"),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case struct",
			args: args{
				any: struct{}{},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Uint32(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Uint32(%v)", tt.args.any)
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case nil",
			args: args{
				any: nil,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case true",
			args: args{
				any: true,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "case false",
			args: args{
				any: false,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "case int",
			args: args{
				any: int(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int8",
			args: args{
				any: int8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int16",
			args: args{
				any: int16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int32",
			args: args{
				any: int32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case int64",
			args: args{
				any: int64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint",
			args: args{
				any: uint(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint8",
			args: args{
				any: uint8(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint16",
			args: args{
				any: uint16(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint32",
			args: args{
				any: uint32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case uint64",
			args: args{
				any: uint64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float32",
			args: args{
				any: float32(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case float64",
			args: args{
				any: float64(8),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case string",
			args: args{
				any: "8",
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case []byte",
			args: args{
				any: []byte("8"),
			},
			want:    8,
			wantErr: assert.NoError,
		},
		{
			name: "case struct",
			args: args{
				any: struct{}{},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Uint64(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Uint64(%v)", tt.args.any)
		})
	}
}
