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
			name: "case 1",
			args: args{
				any: nil,
			},
			want: 0,
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
