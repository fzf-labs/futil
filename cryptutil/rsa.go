package cryptutil

import (
	"github.com/golang-module/dongle"
	"github.com/golang-module/dongle/openssl"
)

// defines key format enum type.
// 定义密钥格式枚举类型
type keyFormat string

// key format constants.
// 密钥格式常量
const (
	PKCS1 keyFormat = "pkcs1"
	PKCS8 keyFormat = "pkcs8"
)

func GenRsaKey(keyFormat keyFormat) (publicKey, privateKey []byte) {
	switch keyFormat {
	case PKCS1:
		return openssl.RSA.GenKeyPair(openssl.PKCS1, 1024)
	case PKCS8:
		return openssl.RSA.GenKeyPair(openssl.PKCS8, 2048)
	}
	return
}

func RsaEncrypt(data string, key []byte) string {
	return dongle.Encrypt.FromString(data).ByRsa(key).ToHexString()
}

func RsaDecrypt(ciphertext string, key []byte) string {
	return dongle.Decrypt.FromHexString(ciphertext).ByRsa(key).ToString()
}

// FormatPrivateKey 格式化 普通应用秘钥
func FormatPrivateKey(keyFormat keyFormat, privateKey []byte) (key []byte) {
	switch keyFormat {
	case PKCS1:
		return openssl.RSA.FormatPrivateKey(openssl.PKCS1, privateKey)
	case PKCS8:
		return openssl.RSA.FormatPrivateKey(openssl.PKCS8, privateKey)
	}
	return
}

// FormatPublicKey 格式化 普通应用公钥
func FormatPublicKey(keyFormat keyFormat, privateKey []byte) (key []byte) {
	switch keyFormat {
	case PKCS1:
		return openssl.RSA.FormatPublicKey(openssl.PKCS1, privateKey)
	case PKCS8:
		return openssl.RSA.FormatPublicKey(openssl.PKCS8, privateKey)
	}
	return
}
