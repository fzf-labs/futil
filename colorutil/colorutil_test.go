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
		wantErr   bool
	}{
		{
			name: "test1",
			args: args{
				colorHex: "#845050",
			},
			wantRed:   132,
			wantGreen: 80,
			wantBlue:  80,
			wantErr:   false,
		},
		{
			name: "test2",
			args: args{
				colorHex: "",
			},
			wantRed:   0,
			wantGreen: 0,
			wantBlue:  0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRed, gotGreen, gotBlue, err := ColorHexToRGB(tt.args.colorHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("ColorHexToRGB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
		{
			name: "test2",
			args: args{
				red:   1,
				green: 1,
				blue:  1,
			},
			want: "#010101",
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
