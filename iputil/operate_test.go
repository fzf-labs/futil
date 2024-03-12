package iputil

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPublicIP(t *testing.T) {
	ip := GetPublicIP()
	assert.True(t, ip != "")
}

func TestGetPublicIPByHTTP(t *testing.T) {
	ip, err := GetPublicIPByHTTP()
	if err != nil {
		return
	}
	assert.True(t, ip != "")
	assert.Equal(t, nil, err)
}

func TestGetRealIP(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				r: &http.Request{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetRealIP(tt.args.r), "GetRealIP(%v)", tt.args.r)
		})
	}
}

func TestIPToLong(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				ip: "127.0.0.1",
			},
			want:    0x7f000001,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IPToLong(tt.args.ip)
			if !tt.wantErr(t, err, fmt.Sprintf("IPToLong(%v)", tt.args.ip)) {
				return
			}
			assert.Equalf(t, tt.want, got, "IPToLong(%v)", tt.args.ip)
		})
	}
}

func TestLongToIP(t *testing.T) {
	type args struct {
		long uint
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
				long: 0x7f000001,
			},
			want:    "127.0.0.1",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LongToIP(tt.args.long)
			if !tt.wantErr(t, err, fmt.Sprintf("LongToIP(%v)", tt.args.long)) {
				return
			}
			assert.Equalf(t, tt.want, got, "LongToIP(%v)", tt.args.long)
		})
	}
}
