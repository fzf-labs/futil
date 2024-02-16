package compressutil

import (
	azip "archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zip"
)

// Zip 压缩
func Zip(srcPath, destZip string) error {
	// 预防：旧文件无法覆盖
	err := os.RemoveAll(destZip)
	if err != nil {
		return err
	}
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer func(zipFile *os.File) {
		_ = zipFile.Close()
	}(zipFile)
	archive := azip.NewWriter(zipFile)
	defer func(archive *azip.Writer) {
		_ = archive.Close()
	}(archive)
	err = filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 获取: 文件头信息
		header, err := azip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(srcPath)+"/")
		// 判断文件是否为文件夹
		if info.IsDir() {
			header.Name += "/"
		} else {
			// 设置: zip 的文件压缩算法
			header.Method = azip.Deflate
		}
		// 创建: 压缩包头部信息
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err2 := os.Open(path)
			if err2 != nil {
				return err2
			}
			defer func(file *os.File) {
				_ = file.Close()
			}(file)
			_, err2 = io.Copy(writer, file)
			if err2 != nil {
				return err2
			}
		}
		return err
	})
	if err != nil {
		return err
	}
	return err
}

// UnZip 解压
func UnZip(srcZip, destPath string) error {
	zipReader, err := zip.OpenReader(srcZip)
	if err != nil {
		return err
	}
	defer func(zipReader *zip.ReadCloser) {
		_ = zipReader.Close()
	}(zipReader)
	for _, f := range zipReader.File {
		path := filepath.Join(destPath, f.Name)
		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}
		if err2 := os.MkdirAll(filepath.Dir(path), os.ModePerm); err2 != nil {
			return err2
		}
		inFile, err2 := f.Open()
		if err2 != nil {
			return err2
		}
		outFile, err2 := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err2 != nil {
			_ = inFile.Close() // 关闭输入文件
			return err2
		}
		_, err2 = io.Copy(outFile, inFile)
		_ = outFile.Close() // 关闭输出文件
		_ = inFile.Close()  // 关闭输入文件
		if err2 != nil {
			return err2
		}
	}
	return nil
}
