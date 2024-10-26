package cryptutil

import (
	"testing"
)

var (
	key = "key"
)

func TestSha1(t *testing.T) {
	data := "test"
	expected := "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
	result := Sha1(data)
	if result != expected {
		t.Errorf("Sha1(%s) = %s; want %s", data, result, expected)
	}
}

func TestSha256(t *testing.T) {
	data := "test"
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	result := Sha256(data)
	if result != expected {
		t.Errorf("Sha256(%s) = %s; want %s", data, result, expected)
	}
}

func TestSha512(t *testing.T) {
	data := "test"
	expected := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	result := Sha512(data)
	if result != expected {
		t.Errorf("Sha512(%s) = %s; want %s", data, result, expected)
	}
}

func TestHmacSha1(t *testing.T) {
	data := "test"

	expected := "671f54ce0c540f78ffe1e26dcf9c2a047aea4fda" // 更新了正确的 HMAC-SHA1 哈希值
	result := HmacSha1(data, key)
	if result != expected {
		t.Errorf("HmacSha1(%s, %s) = %s; want %s", data, key, result, expected)
	}
}

func TestHmacSha256(t *testing.T) {
	data := "test"
	expected := "02afb56304902c656fcb737cdd03de6205bb6d401da2812efd9b2d36a08af159"
	result := HmacSha256(data, key)
	if result != expected {
		t.Errorf("HmacSha256(%s, %s) = %s; want %s", data, key, result, expected)
	}
}

func TestHmacSha512(t *testing.T) {
	data := "test"
	expected := "287a0fb89a7fbdfa5b5538636918e537a5b83065e4ff331268b7aaa115dde047a9b0f4fb5b828608fc0b6327f10055f7637b058e9e0dbb9e698901a3e6dd461c"
	result := HmacSha512(data, key)
	if result != expected {
		t.Errorf("HmacSha512(%s, %s) = %s; want %s", data, key, result, expected)
	}
}
