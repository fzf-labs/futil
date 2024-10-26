package cryptutil

import (
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	publicKey, privateKey := GenRsaKey(PKCS1)
	if len(publicKey) == 0 || len(privateKey) == 0 {
		t.Error("GenRsaKey(PKCS1) failed to generate keys")
	}

	publicKey, privateKey = GenRsaKey(PKCS8)
	if len(publicKey) == 0 || len(privateKey) == 0 {
		t.Error("GenRsaKey(PKCS8) failed to generate keys")
	}
}

func TestRsaEncryptDecrypt(t *testing.T) {
	data := "test"
	publicKey, privateKey := GenRsaKey(PKCS1)

	encrypted := RsaEncrypt(data, publicKey)
	decrypted := RsaDecrypt(encrypted, privateKey)

	if decrypted != data {
		t.Errorf("RsaDecrypt(RsaEncrypt(%s)) = %s; want %s", data, decrypted, data)
	}
}

func TestFormatPrivateKey(t *testing.T) {
	_, privateKey := GenRsaKey(PKCS1)
	formattedKey := FormatPrivateKey(PKCS1, privateKey)
	if len(formattedKey) == 0 {
		t.Error("FormatPrivateKey(PKCS1) failed to format key")
	}

	formattedKey = FormatPrivateKey(PKCS8, privateKey)
	if len(formattedKey) == 0 {
		t.Error("FormatPrivateKey(PKCS8) failed to format key")
	}
}

func TestFormatPublicKey(t *testing.T) {
	publicKey, _ := GenRsaKey(PKCS1)
	formattedKey := FormatPublicKey(PKCS1, publicKey)
	if len(formattedKey) == 0 {
		t.Error("FormatPublicKey(PKCS1) failed to format key")
	}

	formattedKey = FormatPublicKey(PKCS8, publicKey)
	if len(formattedKey) == 0 {
		t.Error("FormatPublicKey(PKCS8) failed to format key")
	}
}
