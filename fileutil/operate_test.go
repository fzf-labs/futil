package fileutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMkdir(t *testing.T) {
	testDir := "test_dir"
	defer os.RemoveAll(testDir)

	err := Mkdir(testDir)
	if err != nil {
		t.Errorf("Mkdir failed: %v", err)
	}

	if !IsPathExists(testDir) {
		t.Error("Directory was not created")
	}
}

func TestBatchMkDir(t *testing.T) {
	testDirs := []string{"test_dir1", "test_dir2", "test_dir3"}
	defer func() {
		for _, dir := range testDirs {
			os.RemoveAll(dir)
		}
	}()

	err := BatchMkDir(testDirs...)
	if err != nil {
		t.Errorf("BatchMkDir failed: %v", err)
	}

	for _, dir := range testDirs {
		if !IsPathExists(dir) {
			t.Errorf("Directory %s was not created", dir)
		}
	}
}

func TestReadFileToString(t *testing.T) {
	testFile := "test_file.txt"
	testContent := "Hello, World!"
	defer os.Remove(testFile)

	// Create test file
	err := WriteContentCover(testFile, testContent)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test reading
	content, err := ReadFileToString(testFile)
	if err != nil {
		t.Errorf("ReadFileToString failed: %v", err)
	}

	if content != testContent {
		t.Errorf("Expected content %q, got %q", testContent, content)
	}
}

func TestWriteContentCover(t *testing.T) {
	testFile := "test_write.txt"
	testContent := "Test Content"
	defer os.Remove(testFile)

	err := WriteContentCover(testFile, testContent)
	if err != nil {
		t.Errorf("WriteContentCover failed: %v", err)
	}

	content, err := ReadFileToString(testFile)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
	}

	if content != testContent {
		t.Errorf("Expected content %q, got %q", testContent, content)
	}
}

func TestCopyFile(t *testing.T) {
	srcFile := "source.txt"
	dstFile := "destination.txt"
	testContent := "Test Copy Content"
	defer func() {
		os.Remove(srcFile)
		os.Remove(dstFile)
	}()

	// Create source file
	err := WriteContentCover(srcFile, testContent)
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Test copying
	err = CopyFile(srcFile, dstFile)
	if err != nil {
		t.Errorf("CopyFile failed: %v", err)
	}

	// Verify content
	dstContent, err := ReadFileToString(dstFile)
	if err != nil {
		t.Errorf("Failed to read destination file: %v", err)
	}

	if dstContent != testContent {
		t.Errorf("Expected content %q, got %q", testContent, dstContent)
	}
}

func TestReadAllFileToMap(t *testing.T) {
	testDir := "test_dir_structure"
	defer os.RemoveAll(testDir)

	// Create test directory structure
	dirs := []string{
		filepath.Join(testDir, "dir1"),
		filepath.Join(testDir, "dir2"),
	}
	files := []string{
		filepath.Join(testDir, "file1.txt"),
		filepath.Join(testDir, "dir1", "file2.txt"),
	}

	err := BatchMkDir(dirs...)
	if err != nil {
		t.Fatalf("Failed to create test directories: %v", err)
	}

	for _, file := range files {
		err = WriteContentCover(file, "test content")
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", file, err)
		}
	}

	// Test reading files
	fileMap, err := ReadAllFileToMap(testDir)
	if err != nil {
		t.Errorf("ReadAllFileToMap failed: %v", err)
	}

	// Verify file count
	expectedCount := len(files) + len(dirs)
	if len(fileMap) != expectedCount {
		t.Errorf("Expected %d files/directories, got %d", expectedCount, len(fileMap))
	}
}

func TestExt(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"test.txt", ".txt"},
		{"test.TXT", ".txt"},
		{"test.jpg", ".jpg"},
		{"test", ""},
		{"test.tar.gz", ".gz"},
	}

	for _, tt := range tests {
		result := Ext(tt.path)
		if result != tt.expected {
			t.Errorf("Ext(%q) = %q, want %q", tt.path, result, tt.expected)
		}
	}
}

func TestFilePrefix(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"test.txt", "test"},
		{"/path/to/file.jpg", "file"},
		{"document", "document"},
		{"path/test.tar.gz", "test.tar"},
	}

	for _, tt := range tests {
		result := FilePrefix(tt.path)
		if result != tt.expected {
			t.Errorf("FilePrefix(%q) = %q, want %q", tt.path, result, tt.expected)
		}
	}
}

func TestRemoveExt(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"test.txt", "test"},
		{"test.TXT", "test"},
		{"test", "test"},
		{"test.tar.gz", "test.tar"},
	}

	for _, tt := range tests {
		result := RemoveExt(tt.path)
		if result != tt.expected {
			t.Errorf("RemoveExt(%q) = %q, want %q", tt.path, result, tt.expected)
		}
	}
}
