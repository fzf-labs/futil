package cryptutil

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Simple string",
			input:    "Hello, World!",
			expected: "SGVsbG8sIFdvcmxkIQ==",
		},
		{
			name:     "String with special characters",
			input:    "Base64 编码测试",
			expected: "QmFzZTY0IOe8lueggea1i+ivlQ==",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Base64Encode(tt.input)
			if result != tt.expected {
				t.Errorf("Base64Encode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Simple string",
			input:    "SGVsbG8sIFdvcmxkIQ==",
			expected: "Hello, World!",
		},
		{
			name:     "String with special characters",
			input:    "QmFzZTY0IOe8lueggea1i+ivlQ==",
			expected: "Base64 编码测试",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Base64Decode(tt.input)
			if result != tt.expected {
				t.Errorf("Base64Decode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBase64EncodeDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Empty string",
			input: "",
		},
		{
			name:  "Simple string",
			input: "Hello, World!",
		},
		{
			name:  "String with special characters",
			input: "Base64 编码测试 !@#$%^&*()_+",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := Base64Encode(tt.input)
			decoded := Base64Decode(encoded)
			if decoded != tt.input {
				t.Errorf("Base64Encode then Base64Decode(%q) = %q, want %q", tt.input, decoded, tt.input)
			}
		})
	}
}
