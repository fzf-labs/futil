package fileutil

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

var (
	// DefaultDirPerm 文件权限
	DefaultDirPerm   os.FileMode = 0775
	DefaultFilePerm  os.FileMode = 0665
	OnlyReadFilePerm os.FileMode = 0444

	// DefaultFileFlags 创建文件 只写 追加
	DefaultFileFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	CoverFileFlags   = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	// OnlyReadFileFlags 只读
	OnlyReadFileFlags = os.O_RDONLY
)

// ************************************************************
//	dir
// ************************************************************

// Mkdir 创建文件夹
func Mkdir(dirPath string) error {
	return os.MkdirAll(dirPath, DefaultDirPerm)
}

// BatchMkDir 批量创建文件夹
func BatchMkDir(dirs ...string) error {
	for _, v := range dirs {
		if !IsPathExists(v) {
			err := os.MkdirAll(v, os.ModePerm)
			if err != nil {
				return err
			}
			err = os.Chmod(v, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// PrintDirAll 打印目录
// example PrintDirAll("/Users/why/Desktop/go/test", 0)
func PrintDirAll(p string, curHer int) {
	fileInfos, err := os.ReadDir(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHer := curHer; tmpHer > 0; tmpHer-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			PrintDirAll(p+"/"+info.Name(), curHer+1)
		} else {
			for tmpHier := curHer; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}

// ReadAllFileToMap 读取所有的文件形成一个map
func ReadAllFileToMap(p string) (map[string]FileInfo, error) {
	infos := make(map[string]FileInfo)
	err := newReadAllFileInfo().doFile(p, infos)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

// ReadAllFileToSli 读取所有的文件形成一个切片
func ReadAllFileToSli(p string) ([]FileInfo, error) {
	res := make([]FileInfo, 0)
	readFileToMap, err := ReadAllFileToMap(p)
	if err != nil {
		return nil, err
	}
	for _, v := range readFileToMap {
		res = append(res, v)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	return res, nil
}

// ReadAllDirToMap 读取所有的文件夹形成一个map
func ReadAllDirToMap(p string) (map[string]FileInfo, error) {
	infos := make(map[string]FileInfo)
	err := newReadAllFileInfo().doDir(p, infos)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

// ReadAllDirToSli 读取所有的文件形成一个切片
func ReadAllDirToSli(p string) ([]FileInfo, error) {
	res := make([]FileInfo, 0)
	readFileToMap, err := ReadAllDirToMap(p)
	if err != nil {
		return nil, err
	}
	for _, v := range readFileToMap {
		res = append(res, v)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	return res, nil
}

type FileInfo struct {
	ID      int64  `json:"id"`
	Pid     int64  `json:"pid"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	File    string `json:"file"`
	IsDir   bool   `json:"isDir"`
	ModTime int64  `json:"modTime"`
	Size    int64  `json:"size"`
}

type readAllFile struct {
	id int64
}

func newReadAllFileInfo() *readAllFile {
	return &readAllFile{id: 0}
}

// 所有的文件
func (r *readAllFile) doFile(p string, files map[string]FileInfo) error {
	pid := r.id
	fileInfos, err := os.ReadDir(p)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		r.id++
		fileInfo, err := info.Info()
		if err != nil {
			return err
		}
		fileName := filepath.Join(p, fileInfo.Name())
		files[fileName] = FileInfo{
			ID:      r.id,
			Pid:     pid,
			Name:    fileInfo.Name(),
			Path:    p,
			File:    fileName,
			IsDir:   fileInfo.IsDir(),
			ModTime: fileInfo.ModTime().Unix(),
			Size:    fileInfo.Size(),
		}
		if info.IsDir() {
			err := r.doFile(filepath.Join(p, info.Name()), files)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 文件夹
func (r *readAllFile) doDir(p string, files map[string]FileInfo) error {
	pid := r.id
	fileInfos, err := os.ReadDir(p)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		r.id++
		fileInfo, err := info.Info()
		if err != nil {
			return err
		}
		if info.IsDir() {
			fileName := filepath.Join(p, fileInfo.Name())
			files[fileName] = FileInfo{
				ID:      r.id,
				Pid:     pid,
				Name:    fileInfo.Name(),
				Path:    p,
				File:    fileName,
				IsDir:   fileInfo.IsDir(),
				ModTime: fileInfo.ModTime().Unix(),
				Size:    fileInfo.Size(),
			}
		}
		if info.IsDir() {
			err := r.doDir(filepath.Join(p, info.Name()), files)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type DeepFileInfo struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	File    string `json:"doFile"`
	IsDir   bool   `json:"isDir"`
	ModTime int64  `json:"modTime"`
	Size    int64  `json:"size"`
}

// ReadDeepFile 读取指定深度的文件
func ReadDeepFile(p string, deep int) (map[string]DeepFileInfo, error) {
	infos := make(map[string]DeepFileInfo, 0)
	err := readDeepFile(p, 0, deep, infos)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func readDeepFile(p string, deepNow, deep int, files map[string]DeepFileInfo) error {
	if deepNow > deep {
		return nil
	}
	if !IsPathExists(p) {
		return nil
	}
	fileInfos, err := os.ReadDir(p)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		fileInfo, err := info.Info()
		if err != nil {
			return err
		}
		if deepNow == deep {
			fileName := filepath.Join(p, fileInfo.Name())
			files[fileName] = DeepFileInfo{
				Name:    fileInfo.Name(),
				Path:    p,
				File:    fileName,
				IsDir:   fileInfo.IsDir(),
				ModTime: fileInfo.ModTime().Unix(),
				Size:    fileInfo.Size(),
			}
		}
		if info.IsDir() {
			err := readDeepFile(filepath.Join(p, info.Name()), deepNow+1, deep, files)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ************************************************************
// files
// ************************************************************

// OpenFile 打开文件，但会自动创建目录。
func OpenFile(fp string, flag int, perm os.FileMode) (*os.File, error) {
	fileDir := filepath.Dir(fp)
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(fp, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// QuickOpenFile 快速打开文件，目录不存在则会自动创建目录。
func QuickOpenFile(filePath string) (*os.File, error) {
	return OpenFile(filePath, DefaultFileFlags, DefaultFilePerm)
}

// OpenReadFile 只读方式打开文件
func OpenReadFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, OnlyReadFileFlags, OnlyReadFilePerm)
}

// ReadFileLineToSli  按行读取文件
func ReadFileLineToSli(dir string) ([]string, error) {
	file, err := os.OpenFile(dir, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadURLFileLineToSli 按行读取url文件
func ReadURLFileLineToSli(url string) ([]string, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadFileToString 读取文件到string
func ReadFileToString(dir string) (string, error) {
	file, err := os.OpenFile(dir, os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	filesize := fileInfo.Size()
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

// ReadFileByURLToByte 读取url中的文件,并转为[]byte格式
func ReadFileByURLToByte(url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

// ************************************************************
//	write, copy files
// ************************************************************

// WriteContentCover 数据写入，不存在则创建
func WriteContentCover(filePath, content string) error {
	fileDir := filepath.Dir(filePath)
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return err
	}
	dstFile, err := os.OpenFile(filePath, CoverFileFlags, DefaultFilePerm)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = dstFile.WriteString(content)
	if err != nil {
		return err
	}
	return err
}

// WriteContentAppend 数据写入，不存在则创建
func WriteContentAppend(filePath, content string) error {
	fileDir := filepath.Dir(filePath)
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return err
	}
	dstFile, err := os.OpenFile(filePath, DefaultFileFlags, DefaultFilePerm)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = dstFile.WriteString(content)
	if err != nil {
		return err
	}
	return err
}

// WriteCsvCover 数据覆盖写入，不存在则创建
func WriteCsvCover(filePath string, content []string) error {
	fileDir := filepath.Dir(filePath)
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return err
	}
	f, err := os.OpenFile(filePath, CoverFileFlags, DefaultFilePerm)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	writer := csv.NewWriter(f)
	for _, v := range content {
		err = writer.Write([]string{v})
		if err != nil {
			return err
		}
	}
	// 将缓存中的内容写入到文件里
	writer.Flush()
	err = writer.Error()
	if err != nil {
		return err
	}
	return nil
}

// WriteCsvDoubleSliCover 数据覆盖写入，不存在则创建
func WriteCsvDoubleSliCover(filePath string, content [][]string) error {
	fileDir := filepath.Dir(filePath)
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return err
	}
	f, err := os.OpenFile(filePath, CoverFileFlags, DefaultFilePerm)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	writer := csv.NewWriter(f)
	err = writer.WriteAll(content)
	if err != nil {
		return err
	}
	// 将缓存中的内容写入到文件里
	writer.Flush()
	err = writer.Error()
	if err != nil {
		return err
	}
	return nil
}

// CopyFile 复制文件
func CopyFile(srcPath, dstPath string) error {
	srcFile, err := os.OpenFile(srcPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// create and open file
	dstFile, err := QuickOpenFile(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// ************************************************************
//	rename
// ************************************************************

// Rename 重命名
func Rename(src, dst string) error {
	return os.Rename(src, dst)
}

// ************************************************************
//	remove
// ************************************************************

// Remove 删除命名文件或 (空) 目录。
func Remove(fPath string) error {
	if IsPathExists(fPath) {
		return os.Remove(fPath)
	}
	return nil
}
func RemoveExt(p string) string {
	ext := filepath.Ext(p)
	if ext == "" {
		return p
	}
	return strings.TrimSuffix(p, ext)
}

// ************************************************************
//	other operates
// ************************************************************

// DownloadFile 会将url下载到本地文件，它会在下载时写入，而不是将整个文件加载到内存中。
func DownloadFile(url, filePath string) error {
	// Get the data
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func FilePrefix(filename string) string {
	name := path.Base(filename)
	return name[0 : len(name)-len(path.Ext(filename))]
}

// Move 移动文件
func Move(srcPath, dstPath string) error {
	err := os.Rename(srcPath, dstPath)
	if err != nil {
		return err
	}
	return nil
}

// Ext 文件扩展名
func Ext(p string) string {
	return strings.ToLower(filepath.Ext(p))
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
