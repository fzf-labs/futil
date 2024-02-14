package conv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByte(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				any: 'a',
			},
			want:    0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Byte(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Byte(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Byte(%v)", tt.args.any)
		})
	}
}

func TestBytes(t *testing.T) {
	type args struct {
		value any
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
			got, err := Bytes(tt.args.value)
			if !tt.wantErr(t, err, fmt.Sprintf("Bytes(%v)", tt.args.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Bytes(%v)", tt.args.value)
		})
	}
}
