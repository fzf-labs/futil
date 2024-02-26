package validutil

import "testing"

func TestIsZero(t *testing.T) {
	type args struct {
		any any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case nil",
			args: args{
				any: nil,
			},
			want: true,
		},
		{
			name: "case string true",
			args: args{
				any: "",
			},
			want: true,
		},
		{
			name: "case string false",
			args: args{
				any: "8",
			},
			want: false,
		},
		{
			name: "case int true",
			args: args{
				any: int(0),
			},
			want: true,
		},
		{
			name: "case int false",
			args: args{
				any: int(1),
			},
			want: false,
		},
		{
			name: "case uint true",
			args: args{
				any: uint(0),
			},
			want: true,
		},
		{
			name: "case uint false",
			args: args{
				any: uint(1),
			},
			want: false,
		},
		{
			name: "case slice true",
			args: args{
				any: []string{},
			},
			want: true,
		},
		{
			name: "case slice false",
			args: args{
				any: []string{"8"},
			},
			want: false,
		},
		{
			name: "case float64 true",
			args: args{
				any: float64(0),
			},
			want: true,
		},
		{
			name: "case float64 false",
			args: args{
				any: float64(8),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.args.any); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIP(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				ipStr: "127.250.255.254",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				ipStr: "123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIP(tt.args.ipStr); got != tt.want {
				t.Errorf("IsIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPV4(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				ipStr: "127.250.255.254",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				ipStr: "123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPV4(tt.args.ipStr); got != tt.want {
				t.Errorf("IsIPV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPV6(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				ipStr: "FC00:0000:130F:0000:0000:09C0:876A:130B",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				ipStr: "123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPV6(tt.args.ipStr); got != tt.want {
				t.Errorf("IsIPV6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPort(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				port: 3306,
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				port: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPort(tt.args.port); got != tt.want {
				t.Errorf("IsPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "www.baidu.com",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "123456",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsURL(tt.args.str); got != tt.want {
				t.Errorf("IsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				email: "123456@qq.com",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				email: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.email); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhoneTight(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				phone: "18888888888",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				phone: "12888888888",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhoneTight(tt.args.phone); got != tt.want {
				t.Errorf("IsPhoneTight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhoneLoose(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				phone: "18888888888",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				phone: "12888888888",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhoneLoose(tt.args.phone); got != tt.want {
				t.Errorf("IsPhoneLoose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTelephone(t *testing.T) {
	type args struct {
		telephone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				telephone: "07552000000",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				telephone: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTelephone(tt.args.telephone); got != tt.want {
				t.Errorf("IsTelephone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPostalCode(t *testing.T) {
	type args struct {
		postalCode string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				postalCode: "518000",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				postalCode: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPostalCode(tt.args.postalCode); got != tt.want {
				t.Errorf("IsPostalCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsResidentID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				id: "230101198910059687",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				id: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsResidentID(tt.args.id); got != tt.want {
				t.Errorf("IsResidentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsQQ(t *testing.T) {
	type args struct {
		qq string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				qq: "123456",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsQQ(tt.args.qq); got != tt.want {
				t.Errorf("IsQQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPassport(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				p: "E12345678",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				p: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPassport(tt.args.p); got != tt.want {
				t.Errorf("IsPassport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeakPassword(t *testing.T) {
	type args struct {
		password string
		length   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				password: "123456",
				length:   6,
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				password: "123456789+",
				length:   6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWeakPassword(tt.args.password, tt.args.length); got != tt.want {
				t.Errorf("IsWeakPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStrongPassword(t *testing.T) {
	type args struct {
		password string
		length   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				password: "Abc123!@#",
				length:   6,
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				password: "123456",
				length:   6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStrongPassword(tt.args.password, tt.args.length); got != tt.want {
				t.Errorf("IsStrongPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDomain(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				p: "www.baidu.com",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				p: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDomain(tt.args.p); got != tt.want {
				t.Errorf("IsDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMac(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				p: "5E-05-83-86-46-EF",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				p: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMac(tt.args.p); got != tt.want {
				t.Errorf("IsMac() = %v, want %v", got, tt.want)
			}
		})
	}
}
