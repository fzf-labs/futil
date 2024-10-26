//nolint:all
package cryptutil

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/golang-module/dongle"
)

// Md5ToHex 返回字符串的 md5 值，并转换为十六进制字符串
func Md5ToHex(data string) string {
	return dongle.Encrypt.FromString(data).ByMd5().ToHexString()
}

// Md5ToBase64 返回字符串的 md5 值，并转换为 base64 字符串
func Md5ToBase64(data string) string {
	return dongle.Encrypt.FromString(data).ByMd5().ToBase64String()
}

// Md5File 返回文件的 md5 值
func Md5File(filename string) (string, error) {
	if fileInfo, err := os.Stat(filename); err != nil {
		return "", err
	} else if fileInfo.IsDir() {
		return "", nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	chunkSize := 65536
	for buf, reader := make([]byte, chunkSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hash.Write(buf[:n])
	}
	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}
