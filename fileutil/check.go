package fileutil

import (
	"bytes"
	"os"
	"path"
)

// IsPathExists 检查文件或目录是否存在。
func IsPathExists(fp string) bool {
	if fp == "" {
		return false
	}

	if _, err := os.Stat(fp); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsDir 检查是否是目录
func IsDir(fp string) bool {
	if fp == "" {
		return false
	}

	if fi, err := os.Stat(fp); err == nil {
		return fi.IsDir()
	}
	return false
}

// IsFile 检查是否是文件
func IsFile(fp string) bool {
	if fp == "" {
		return false
	}

	if fi, err := os.Stat(fp); err == nil {
		return !fi.IsDir()
	}
	return false
}

// IsAbsPath 是否是相对路径
func IsAbsPath(fp string) bool {
	return path.IsAbs(fp)
}

// IsImageFile 检查文件是图像文件
func IsImageFile(fp string) bool {
	mime := MimeType(fp)
	if mime == "" {
		return false
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

	for _, imgMime := range ImageMimeTypes {
		if imgMime == mime {
			return true
		}
	}
	return false
}

// IsZipFile 检查是zip文件
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func IsZipFile(fp string) bool {
	f, err := os.Open(fp)
	if err != nil {
		return false
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}
	return bytes.Equal(buf, []byte("PK\x03\x04"))
}
