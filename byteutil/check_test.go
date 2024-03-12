package byteutil

import "testing"

func TestIsLetterUpper(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				b: 'a',
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				b: 'A',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetterUpper(tt.args.b); got != tt.want {
				t.Errorf("IsLetterUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLetterLower(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				b: 'A',
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				b: 'a',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetterLower(tt.args.b); got != tt.want {
				t.Errorf("IsLetterLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				b: 'A',
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				b: 'a',
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				b: '0',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetter(tt.args.b); got != tt.want {
				t.Errorf("IsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumChar(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				b: 'A',
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				b: 'a',
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				b: '0',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumChar(tt.args.b); got != tt.want {
				t.Errorf("IsNumChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				b: 'A',
			},
			want: true,
		},
		{
			name: "case true",
			args: args{
				b: 'a',
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				b: '0',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpha(tt.args.b); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}
