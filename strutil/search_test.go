package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrPos(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s:   "abc",
				sub: "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StrPos(tt.args.s, tt.args.sub), "StrPos(%v, %v)", tt.args.s, tt.args.sub)
		})
	}
}
