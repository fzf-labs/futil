package conv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat32(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    float32
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
			got, err := Float32(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Float32(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Float32(%v)", tt.args.any)
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    float64
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
			got, err := Float64(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Float64(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Float64(%v)", tt.args.any)
		})
	}
}
