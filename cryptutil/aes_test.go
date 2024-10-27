package cryptutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cipher, _ = NewAesCipher("CBC", "PKCS7", "1234567890123456", "1234567890123456")

func TestAesEncrypt(t *testing.T) {
	encrypt := AesEncrypt(cipher, "18888888888")
	assert.NotEmpty(t, encrypt)
}

func TestAesDecrypt(t *testing.T) {
	encrypt := AesEncrypt(cipher, "18888888888")
	assert.NotEmpty(t, AesDecrypt(cipher, encrypt))
}
