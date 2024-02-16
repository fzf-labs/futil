package compressutil

import "testing"

func TestZip(t *testing.T) {
	type args struct {
		srcPath string
		destZip string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				srcPath: "./",
				destZip: "./test.zip",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Zip(tt.args.srcPath, tt.args.destZip); (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnZip(t *testing.T) {
	type args struct {
		srcZip   string
		destPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				srcZip:   "./test.zip",
				destPath: "./test/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnZip(tt.args.srcZip, tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("UnZip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
