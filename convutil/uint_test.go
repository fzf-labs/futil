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
			name:    "",
			args:    args{},
			want:    0,
			wantErr: nil,
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
			name: "",
			args: args{
				any: 'a',
			},
			want:    0,
			wantErr: assert.NoError,
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
			name:    "",
			args:    args{},
			want:    0,
			wantErr: nil,
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
			name:    "",
			args:    args{},
			want:    0,
			wantErr: nil,
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
			name:    "",
			args:    args{},
			want:    0,
			wantErr: nil,
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
