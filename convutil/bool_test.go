package conv

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	tInt int = 1
)

func TestBool(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case nil",
			args: args{
				any: nil,
			},
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name: "case int",
			args: args{
				any: int(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case int8",
			args: args{
				any: int8(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case int16",
			args: args{
				any: int16(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case int32",
			args: args{
				any: int32(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case int64",
			args: args{
				any: int64(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case uint",
			args: args{
				any: uint(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case uint8",
			args: args{
				any: uint8(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case uint16",
			args: args{
				any: uint16(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case uint32",
			args: args{
				any: uint32(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case uint64",
			args: args{
				any: uint64(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case float32",
			args: args{
				any: float32(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case float64",
			args: args{
				any: float64(8),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case bool",
			args: args{
				any: true,
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case string",
			args: args{
				any: string("8"),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case bytes",
			args: args{
				any: []byte("8"),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case time",
			args: args{
				any: time.Now(),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case time",
			args: args{
				any: time.Second,
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case json.Number",
			args: args{
				any: json.Number(""),
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case 20",
			args: args{
				any: "no",
			},
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name: "case 21",
			args: args{
				any: []string{"1", "2"},
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case 22",
			args: args{
				any: [2]string{"1", "2"},
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case 23",
			args: args{
				any: map[string]string{"1": "1", "2": "2"},
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case 24",
			args: args{
				any: &tInt,
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case 25",
			args: args{
				any: struct{}{},
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "case 26",
			args: args{
				any: "no",
			},
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name: "case 26",
			args: args{
				any: []byte("no"),
			},
			want:    false,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bool(tt.args.any)
			if !tt.wantErr(t, err, fmt.Sprintf("Bool(%v)", tt.args.any)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Bool(%v)", tt.args.any)
		})
	}
}
