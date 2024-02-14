package conv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRune(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name    string
		args    args
		want    rune
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
			got, err := Rune(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Rune(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Rune(%v)", tt.args.any)
		})
	}
}

func TestRunes(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Runes(tt.args.any), "Runes(%v)", tt.args.any)
		})
	}
}
