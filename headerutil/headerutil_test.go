package headerutil

import "testing"

func TestAcceptLanguage(t *testing.T) {
	type args struct {
		acceptLanguage string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				acceptLanguage: "fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5",
			},
			want: "fr-CH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AcceptLanguage(tt.args.acceptLanguage); got != tt.want {
				t.Errorf("AcceptLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}
