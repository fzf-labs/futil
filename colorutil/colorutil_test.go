package colorutil

import "testing"

func TestColorHexToRGB(t *testing.T) {
	type args struct {
		colorHex string
	}
	tests := []struct {
		name      string
		args      args
		wantRed   int
		wantGreen int
		wantBlue  int
	}{
		{
			name: "test1",
			args: args{
				colorHex: "#845050",
			},
			wantRed:   132,
			wantGreen: 80,
			wantBlue:  80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRed, gotGreen, gotBlue := ColorHexToRGB(tt.args.colorHex)
			if gotRed != tt.wantRed {
				t.Errorf("ColorHexToRGB() gotRed = %v, want %v", gotRed, tt.wantRed)
			}
			if gotGreen != tt.wantGreen {
				t.Errorf("ColorHexToRGB() gotGreen = %v, want %v", gotGreen, tt.wantGreen)
			}
			if gotBlue != tt.wantBlue {
				t.Errorf("ColorHexToRGB() gotBlue = %v, want %v", gotBlue, tt.wantBlue)
			}
		})
	}
}

func TestColorRGBToHex(t *testing.T) {
	type args struct {
		red   int
		green int
		blue  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				red:   132,
				green: 80,
				blue:  80,
			},
			want: "#845050",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColorRGBToHex(tt.args.red, tt.args.green, tt.args.blue); got != tt.want {
				t.Errorf("ColorRGBToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
