package urlutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "www.baidu.com/?query=golang",
			},
			want: "www.baidu.com/?query%3Dgolang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, URLEncode(tt.args.s), "URLEncode(%v)", tt.args.s)
		})
	}
}

func TestURLDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "www.baidu.com/?query%3Dgolang",
			},
			want: "www.baidu.com/?query=golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, URLDecode(tt.args.s), "URLDecode(%v)", tt.args.s)
		})
	}
}

func TestRawURLEncode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "www.baidu.com/?query=golang",
			},
			want: "www.baidu.com%2F%3Fquery%3Dgolang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RawURLEncode(tt.args.str), "RawURLEncode(%v)", tt.args.str)
		})
	}
}

func TestRawURLDecode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				str: "www.baidu.com%2F%3Fquery%3Dgolang",
			},
			want:    "www.baidu.com/?query=golang",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLDecode(tt.args.str)
			if !tt.wantErr(t, err, fmt.Sprintf("RawURLDecode(%v)", tt.args.str)) {
				return
			}
			assert.Equalf(t, tt.want, got, "RawURLDecode(%v)", tt.args.str)
		})
	}
}

func TestURLEncodeByMap(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				m: map[string]string{
					"a": "1",
					"b": "2",
				},
			},
			want: "a=1&b=2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, URLEncodeByMap(tt.args.m), "URLEncodeByMap(%v)", tt.args.m)
		})
	}
}

func TestURLDecodeToMap(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "case 1",
			args: args{
				str: "a=1&b=2",
			},
			want: map[string]string{
				"a": "1",
				"b": "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, URLDecodeToMap(tt.args.str), "URLDecodeToMap(%v)", tt.args.str)
		})
	}
}
