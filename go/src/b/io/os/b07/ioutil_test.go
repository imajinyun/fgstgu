package b07_test

import (
	"bytes"
	. "io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func checkSize(t *testing.T, path string, size int64) {
	dir, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat %q (looking for size %d): %s", path, size, err)
	}
	if dir.Size() != size {
		t.Errorf("Stat %q: size %d want %d", path, dir.Size(), size)
	}
}

func TestReadFile(t *testing.T) {
	filename := "TestReadFile"
	contents, err := ReadFile(filename)
	if err == nil {
		t.Fatalf("ReadFile %s: error expected, none found", filename)
	}

	filename = "ioutil_test.go"
	contents, err = ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", filename, err)
	}

	checkSize(t, filename, int64(len(contents)))
}

func TestWriteFile(t *testing.T) {
	f, err := TempFile("", "ioutil-test")
	if err != nil {
		t.Fatal(err)
	}
	filename := f.Name()
	data := "Programming today is a race between software engineers striving to " +
		"build bigger and better idiot-proof programs, and the Universe trying " +
		"to produce bigger and better idiots. So far, the Universe is winning."
	if err := WriteFile(filename, []byte(data), 0644); err != nil {
		t.Fatalf("WriteFile %s: %v", filename, err)
	}

	contents, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", filename, err)
	}
	if string(contents) != data {
		t.Fatalf("contents = %q\nexpected = %q", string(contents), data)
	}
	f.Close()
	os.Remove(filename)
}

func TestReadOnlyWriteFile(t *testing.T) {
	if os.Getuid() == 0 {
		t.Skip("Root can write to read-only files anyway, so skip the read-only test.")
	}

	tempDir, err := TempDir("", t.Name())
	if err != nil {
		t.Fatalf("TempDir %s: %v", t.Name(), err)
	}
	defer os.RemoveAll(tempDir)
	filename := filepath.Join(tempDir, "test.txt")
	hello := []byte("Hello")
	world := []byte("World")
	err = WriteFile(filename, hello, 0444)
	if err != nil {
		t.Fatalf("WriteFile %s: %v", filename, err)
	}
	err = WriteFile(filename, world, 0444)
	if err == nil {
		t.Fatalf("Expected an error when writing to read-only file %s", filename)
	}
	got, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", filename, err)
	}
	if !bytes.Equal(got, hello) {
		t.Fatalf("want %s, got %s", hello, got)
	}
}

func TestReadDir(t *testing.T) {
	dirname := "TestReadDir"
	_, err := ReadDir(dirname)
	if err == nil {
		t.Fatalf("ReadDir %s: error expected, none found", dirname)
	}

	dirname = ".."
	list, err := ReadDir(dirname)
	if err != nil {
		t.Fatalf("ReadDir %s: %v", dirname, err)
	}

	foundFile := false
	foundSubDir := false
	for _, dir := range list {
		switch {
		case !dir.IsDir() && dir.Name() == "data.txt":
			foundFile = true
		case dir.IsDir() && dir.Name() == "b07":
			foundSubDir = true
		}
	}
	if !foundFile {
		t.Fatalf("ReadDir %s: ioutil_test.go file not found", dirname)
	}

	if !foundSubDir {
		t.Fatalf("ReadDir %s: io directory not found", dirname)
	}
}
