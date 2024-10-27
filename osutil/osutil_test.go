package osutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostName(t *testing.T) {
	assert.Equal(t, true, GetHostName() != "")
}

func TestGetOS(t *testing.T) {
	assert.Equal(t, true, GetOS() != "")
}

func TestIsOS(t *testing.T) {
	assert.Equal(t, true, IsWindows() == true || IsMac() == true || IsLinux() == true)
}

func TestGetOsEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				key: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOsEnv(tt.args.key); got != tt.want {
				t.Errorf("GetOsEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetOsEnv(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				key:   "",
				value: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetOsEnv(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SetOsEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveOsEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				key: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveOsEnv(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RemoveOsEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOsBits(t *testing.T) {
	assert.Equal(t, true, GetOsBits() == 32 || GetOsBits() == 64)
}

func TestGetGoroutineID(t *testing.T) {
	assert.Equal(t, true, GetOsBits() == 32 || GetOsBits() == 64)
}
