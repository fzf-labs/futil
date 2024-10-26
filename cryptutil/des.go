package cryptutil

import (
	"github.com/golang-module/dongle"
	"github.com/pkg/errors"
)

// NewDesCipher 返回一个新的DES密码实例。
func NewDesCipher(mode, padding, key, iv string) (*dongle.Cipher, error) {
	cipher := dongle.NewCipher()
	switch mode {
	case "CBC":
		cipher.SetMode(dongle.CBC)
	case "CFB":
		cipher.SetMode(dongle.CFB)
	case "OFB":
		cipher.SetMode(dongle.OFB)
	case "CTR":
		cipher.SetMode(dongle.CTR)
	case "ECB":
		cipher.SetMode(dongle.ECB)
	default:
		return nil, errors.New("invalid mode")
	}
	switch padding {
	case "No":
		cipher.SetPadding(dongle.No)
	case "Empty":
		cipher.SetPadding(dongle.Empty)
	case "Zero":
		cipher.SetPadding(dongle.Zero)
	case "PKCS5":
		cipher.SetPadding(dongle.PKCS5)
	case "PKCS7":
		cipher.SetPadding(dongle.PKCS7)
	case "AnsiX923":
		cipher.SetPadding(dongle.AnsiX923)
	case "ISO97971":
		cipher.SetPadding(dongle.ISO97971)
	default:
		return nil, errors.New("invalid padding")
	}
	if len(key) != 8 {
		return nil, errors.New("invalid key length")
	}
	if len(iv) != 8 {
		return nil, errors.New("invalid iv length")
	}
	cipher.SetKey(key) // key 长度必须是 8 字节
	cipher.SetIV(iv)   // iv 长度必须是 8 字节
	return cipher, nil
}

// DesEncrypt 使用 Des 加密算法加密数据。
func DesEncrypt(cipher *dongle.Cipher, data string) string {
	return dongle.Encrypt.FromString(data).ByDes(cipher).ToRawString()
}

// DesDecrypt 使用 Des 加密算法解密数据。
func DesDecrypt(cipher *dongle.Cipher, data string) string {
	return dongle.Decrypt.FromRawString(data).ByDes(cipher).ToString()
}
