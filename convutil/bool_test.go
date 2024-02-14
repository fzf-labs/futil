package conv

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	tInt        int           = 1
	tInt8       int8          = 1
	tInt16      int16         = 1
	tInt32      int32         = 1
	tInt64      int64         = 1
	tUint       uint          = 1
	tUint8      uint8         = 1
	tUint16     uint16        = 1
	tUint32     uint32        = 1
	tUint64     uint64        = 1
	tFloat32    float32       = 1
	tFloat64    float64       = 1
	tBool       bool          = true
	tString     string        = "1"
	tBytes      []byte        = []byte("1")
	tTime       time.Time     = time.Now()
	tDuration   time.Duration = time.Second
	tJsonNumber               = json.Number("1")
	tStruct     struct{}
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
			name: "case 1",
			args: args{
				any: nil,
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				any: tInt,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				any: tInt8,
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				any: tInt16,
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				any: tInt32,
			},
			want: true,
		},
		{
			name: "case 6",
			args: args{
				any: tInt64,
			},
			want: true,
		},
		{
			name: "case 7",
			args: args{
				any: tUint,
			},
			want: true,
		},
		{
			name: "case 8",
			args: args{
				any: tUint8,
			},
			want: true,
		},
		{
			name: "case 9",
			args: args{
				any: tUint16,
			},
			want: true,
		},
		{
			name: "case 10",
			args: args{
				any: tUint32,
			},
			want: true,
		},
		{
			name: "case 11",
			args: args{
				any: tUint64,
			},
			want: true,
		},
		{
			name: "case 12",
			args: args{
				any: tFloat32,
			},
			want: true,
		},
		{
			name: "case 13",
			args: args{
				any: tFloat64,
			},
			want: true,
		},
		{
			name: "case 14",
			args: args{
				any: tBool,
			},
			want: true,
		},
		{
			name: "case 15",
			args: args{
				any: tString,
			},
			want: true,
		},
		{
			name: "case 16",
			args: args{
				any: tBytes,
			},
			want: true,
		},
		{
			name: "case 17",
			args: args{
				any: tTime,
			},
			want: true,
		},
		{
			name: "case 18",
			args: args{
				any: tDuration,
			},
			want: true,
		},
		{
			name: "case 19",
			args: args{
				any: tJsonNumber,
			},
			want: true,
		},
		{
			name: "case 20",
			args: args{
				any: "no",
			},
			want: false,
		},
		{
			name: "case 21",
			args: args{
				any: []string{"1", "2"},
			},
			want: true,
		},
		{
			name: "case 22",
			args: args{
				any: [2]string{"1", "2"},
			},
			want: true,
		},
		{
			name: "case 23",
			args: args{
				any: map[string]string{"1": "1", "2": "2"},
			},
			want: true,
		},
		{
			name: "case 24",
			args: args{
				any: &tInt,
			},
			want: true,
		},
		{
			name: "case 25",
			args: args{
				any: tStruct,
			},
			want: true,
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
