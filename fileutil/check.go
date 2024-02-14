package fileutil

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gabriel-vasile/mimetype"
)

// PathExists 检查文件或目录是否存在。
func PathExists(p string) bool {
	if p == "" {
		return false
	}

	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsDir 检查是否是目录
func IsDir(p string) bool {
	if p == "" {
		return false
	}

	if fi, err := os.Stat(p); err == nil {
		return fi.IsDir()
	}
	return false
}

// FileExists 检查文件是否存在。
func FileExists(p string) bool {
	return IsFile(p)
}

// IsFile 检查是否是文件。
func IsFile(p string) bool {
	if p == "" {
		return false
	}

	if fi, err := os.Stat(p); err == nil {
		return !fi.IsDir()
	}
	return false
}

// IsAbsPath 是否是相对路径
func IsAbsPath(aPath string) bool {
	return path.IsAbs(aPath)
}

// ImageMimeTypes 图片类型
var ImageMimeTypes = map[string]string{
	"bmp": "image/bmp",
	"gif": "image/gif",
	"ief": "image/ief",
	"jpg": "image/jpeg",
	// "jpe":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"svg":  "image/svg+xml",
	"ico":  "image/x-icon",
	"webp": "image/webp",
}

// IsImageFile check file is image file.
func IsImageFile(p string) bool {
	mime := MimeType(p)
	if mime == "" {
		return false
	}

	for _, imgMime := range ImageMimeTypes {
		if imgMime == mime {
			return true
		}
	}
	return false
}

// IsZipFile check is zip file.
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func IsZipFile(filepath string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// MimeType 获取文件 Mime 类型名称。例如“image/png”
func MimeType(p string) (mime string) {
	file, err := os.Open(p)
	if err != nil {
		return
	}
	return ReaderMimeType(file)
}

// ReaderMimeType 获取文件 Mime 类型名称
func ReaderMimeType(r io.Reader) (mime string) {
	// 512 嗅探长度，用于检测文件 mime 类型
	var buf [512]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return ""
	}
	return http.DetectContentType(buf[:n])
}

// ReaderMimeTypeAndExt 获取文件的mime和ext后缀(使用扩展包)
func ReaderMimeTypeAndExt(r io.Reader) (mime, ext string) {
	reader, err := mimetype.DetectReader(r)
	if err != nil {
		return "", ""
	}
	mime = reader.String()
	ext = reader.Extension()
	return
}
