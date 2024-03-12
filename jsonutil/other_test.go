package jsonutil

import "testing"

func TestStripComments(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				src: "/* comment */",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripComments(tt.args.src); got != tt.want {
				t.Errorf("StripComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
