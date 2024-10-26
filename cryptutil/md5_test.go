package cryptutil

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestMd5Hash(t *testing.T) {
	data := "test"
	expectedHash := "098f6bcd4621d373cade4e832627b4f6"

	hash := md5.New()
	hash.Write([]byte(data))
	hashedData := hex.EncodeToString(hash.Sum(nil))

	if hashedData != expectedHash {
		t.Errorf("Md5Hash(%s) = %s; want %s", data, hashedData, expectedHash)
	}
}

func TestMd5HashEmptyString(t *testing.T) {
	data := ""
	expectedHash := "d41d8cd98f00b204e9800998ecf8427e"

	hash := md5.New()
	hash.Write([]byte(data))
	hashedData := hex.EncodeToString(hash.Sum(nil))

	if hashedData != expectedHash {
		t.Errorf("Md5Hash(%s) = %s; want %s", data, hashedData, expectedHash)
	}
}
