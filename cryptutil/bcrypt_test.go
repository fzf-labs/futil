package cryptutil

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "123456",
			args: args{
				hashedPassword: "$2a$10$5UgbrSkN2sAzUbvf362HReUjXGNE/rdeU0QrDDbb9f87s7CEQGp9a",
				password:       "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Compare(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "123456",
			args: args{password: "123456"},
			want: "$2a$10$wuDyT.MBZqSNtxuBpKvQa.uP.dHMoPRjPNk48Tj8GAqsajt/oPcYS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBcryptSign(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "Simple string",
			data: "Hello, World!",
			want: "$2a$10$",
		},
		{
			name: "Complex string",
			data: "BcryptSign 测试 !@#$%^&*()_+",
			want: "$2a$10$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BcryptSign(tt.data)
			if !strings.HasPrefix(got, tt.want) {
				t.Errorf("BcryptSign() = %v, want prefix %v", got, tt.want)
			}
		})
	}
}

func TestBcryptVerify(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			name: "Simple string",
			data: "Hello, World!",
			want: true,
		},
		{
			name: "Complex string",
			data: "BcryptVerify 测试 !@#$%^&*()_+",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sign := BcryptSign(tt.data)
			if got := BcryptVerify(sign, tt.data); got != tt.want {
				t.Errorf("BcryptVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBcryptVerifyFail(t *testing.T) {
	originalData := "Hello, World!"
	sign := BcryptSign(originalData)

	modifiedData := "Hello, World"
	if got := BcryptVerify(sign, modifiedData); got {
		t.Errorf("BcryptVerify() = %v, want false for modified data", got)
	}
}
