package conv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
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
			got, err := StructToMap(tt.args.value)
			if !tt.wantErr(t, err, fmt.Sprintf("StructToMap(%v)", tt.args.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "StructToMap(%v)", tt.args.value)
		})
	}
}
