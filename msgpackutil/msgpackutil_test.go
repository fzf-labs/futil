package msgpackutil

import (
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    []byte
		wantErr bool
	}{
		{
			name:    "marshal string",
			input:   "hello",
			want:    []byte{165, 104, 101, 108, 108, 111}, // msgpack encoded "hello"
			wantErr: false,
		},
		{
			name:    "marshal int",
			input:   42,
			want:    []byte{42}, // msgpack encoded 42
			wantErr: false,
		},
		{
			name: "marshal map",
			input: map[string]string{
				"key": "value",
			},
			want:    []byte{129, 163, 107, 101, 121, 165, 118, 97, 108, 117, 101}, // msgpack encoded map
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		ptr     interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "unmarshal string",
			input:   []byte{165, 104, 101, 108, 108, 111}, // msgpack encoded "hello"
			ptr:     new(string),
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "unmarshal int",
			input:   []byte{42}, // msgpack encoded 42
			ptr:     new(int),
			want:    42,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Unmarshal(tt.input, tt.ptr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := reflect.ValueOf(tt.ptr).Elem().Interface()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalString(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "marshal string to string",
			input:   "hello",
			want:    string([]byte{165, 104, 101, 108, 108, 111}), // msgpack encoded "hello"
			wantErr: false,
		},
		{
			name:    "marshal int to string",
			input:   42,
			want:    string([]byte{42}), // msgpack encoded 42
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MarshalString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		ptr     interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "unmarshal string from string",
			input:   string([]byte{165, 104, 101, 108, 108, 111}), // msgpack encoded "hello"
			ptr:     new(string),
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "unmarshal int from string",
			input:   string([]byte{42}), // msgpack encoded 42
			ptr:     new(int),
			want:    42,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UnmarshalString(tt.input, tt.ptr)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := reflect.ValueOf(tt.ptr).Elem().Interface()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
