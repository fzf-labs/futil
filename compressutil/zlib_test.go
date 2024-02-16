package compressutil

import (
	"reflect"
	"testing"
)

func TestZlibCompress(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data: []byte("hello world"),
			},
			want:    []byte{120, 156, 0, 11, 0, 244, 255, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 3, 0, 26, 11, 4, 93},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ZlibCompress(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZlibCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZlibCompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZlibUnCompress(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data: []byte{120, 156, 0, 11, 0, 244, 255, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 3, 0, 26, 11, 4, 93},
			},
			want:    []byte("hello world"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ZlibUnCompress(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZlibUnCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZlibUnCompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
