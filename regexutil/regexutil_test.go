package regexutil

import "testing"

func TestIsRegexMatch(t *testing.T) {
	type args struct {
		str   string
		regex string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str:   "18888888888",
				regex: "[1][3,4,5,7,8][0-9]{9}",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str:   "1",
				regex: "/w",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRegexMatch(tt.args.str, tt.args.regex); got != tt.want {
				t.Errorf("IsRegexMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
