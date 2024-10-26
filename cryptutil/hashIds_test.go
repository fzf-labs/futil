package cryptutil

import (
	"testing"
)

func TestHashidsEncodeDecode(t *testing.T) {
	secret := "mysecret"

	length := 8
	hash := NewHashids(secret, length)

	params := []int{1, 2, 3}
	encoded, err := hash.HashidsEncode(params)
	if err != nil {
		t.Fatalf("HashidsEncode failed: %v", err)
	}

	decoded, err := hash.HashidsDecode(encoded)
	if err != nil {
		t.Fatalf("HashidsDecode failed: %v", err)
	}

	if len(decoded) != len(params) {
		t.Fatalf("Decoded length mismatch: got %d, want %d", len(decoded), len(params))
	}

	for i, v := range decoded {
		if v != params[i] {
			t.Errorf("Decoded value mismatch at index %d: got %d, want %d", i, v, params[i])
		}
	}
}

func TestHashidsEncodeError(t *testing.T) {
	secret := "mysecret"
	length := 8
	hash := NewHashids(secret, length)

	// Test with an empty slice
	params := []int{}
	_, err := hash.HashidsEncode(params)
	if err == nil {
		t.Error("Expected error for empty params, got nil")
	}
}

func TestHashidsDecodeError(t *testing.T) {
	secret := "mysecret"
	length := 8
	hash := NewHashids(secret, length)

	// Test with an invalid hash string
	invalidHash := "invalid"
	_, err := hash.HashidsDecode(invalidHash)
	if err == nil {
		t.Error("Expected error for invalid hash, got nil")
	}
}
