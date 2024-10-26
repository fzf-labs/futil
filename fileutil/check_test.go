package fileutil

import "testing"

func TestIsPathExists(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				fp: "./",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				fp: "./check",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPathExists(tt.args.fp); got != tt.want {
				t.Errorf("IsPathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				fp: "./",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(tt.args.fp); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFile(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				fp: "./check.go",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFile(tt.args.fp); got != tt.want {
				t.Errorf("IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAbsPath(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				fp: "./",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAbsPath(tt.args.fp); got != tt.want {
				t.Errorf("IsAbsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsImageFile(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				fp: "./check.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsImageFile(tt.args.fp); got != tt.want {
				t.Errorf("IsImageFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZipFile(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				fp: "./check.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZipFile(tt.args.fp); got != tt.want {
				t.Errorf("IsZipFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
