package cryptutil

import (
	"testing"
)

func TestNewDesCipher(t *testing.T) {
	tests := []struct {
		name    string
		mode    string
		padding string
		key     string
		iv      string
		wantErr bool
	}{
		{
			name:    "Valid CBC mode with PKCS5 padding",
			mode:    "CBC",
			padding: "PKCS5",
			key:     "12345678",
			iv:      "12345678",
			wantErr: false,
		},
		{
			name:    "Invalid mode",
			mode:    "INVALID",
			padding: "PKCS5",
			key:     "12345678",
			iv:      "12345678",
			wantErr: true,
		},
		{
			name:    "Invalid padding",
			mode:    "CBC",
			padding: "INVALID",
			key:     "12345678",
			iv:      "12345678",
			wantErr: true,
		},
		{
			name:    "Invalid key length",
			mode:    "CBC",
			padding: "PKCS5",
			key:     "123",
			iv:      "12345678",
			wantErr: true,
		},
		{
			name:    "Invalid iv length",
			mode:    "CBC",
			padding: "PKCS5",
			key:     "12345678",
			iv:      "123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDesCipher(tt.mode, tt.padding, tt.key, tt.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDesCipher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDesEncryptDecrypt(t *testing.T) {
	cipher, err := NewDesCipher("CBC", "PKCS5", "12345678", "12345678")
	if err != nil {
		t.Fatalf("Failed to create DES cipher: %v", err)
	}

	originalText := "Hello, DES!"
	encryptedText := DesEncrypt(cipher, originalText)
	decryptedText := DesDecrypt(cipher, encryptedText)

	if decryptedText != originalText {
		t.Errorf("DesEncrypt/DesDecrypt failed: got %v, want %v", decryptedText, originalText)
	}
}
