package compressutil

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zlib"
)

// ZlibCompress zlib压缩
func ZlibCompress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	_ = w.Close()
	return b.Bytes(), nil
}

// ZlibUnCompress zlib解压
func ZlibUnCompress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := bytes.NewReader(data)
	r, err := zlib.NewReader(w)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&b, r)
	if err != nil {
		return nil, err
	}
	_ = r.Close()
	return b.Bytes(), nil
}
