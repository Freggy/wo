package driver

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	filePath = "/tmp/wo-fs/file"
	dirPath  = "/tmp/wo-fs/dir"
)

func TestFilesystemDriverRead(t *testing.T) {
	setupFile(t, filePath)
	driver := filesystemDriver{}

	b, err := driver.Read(filePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []byte("test"), b)
}

func TestFilesystemDriverReadDirectory(t *testing.T) {
	setupDir(t, dirPath)
	driver := filesystemDriver{}

	_, err := driver.Read(dirPath)
	if err == nil {
		t.Fatal("error expected")
	}

	assert.EqualError(t, ErrPathIsDir, err.Error())
}

func TestFilesystemDriverWrite(t *testing.T) {
	setupFile(t, filePath)
	driver := filesystemDriver{}
	content := []byte("helloworld")

	if err := driver.Write(filePath, content); err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, content, b)
}

func TestFilesystemDriverWriteDirectory(t *testing.T) {
	setupDir(t, dirPath)
	driver := filesystemDriver{}

	err := driver.Write(dirPath, []byte("helloworld"))
	if err == nil {
		t.Fatal("error expected")
	}

	assert.EqualError(t, ErrPathIsDir, err.Error())
}

func TestFilesystemDriverFileExists(t *testing.T) {
	setupDir(t, filePath)
	driver := filesystemDriver{}
	ok, err := driver.Exists(filePath)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, ok)
}

func TestFilesystemDriverFileNotExist(t *testing.T) {
	setupDir(t, filePath)
	driver := filesystemDriver{}
	ok, err := driver.Exists(dirPath)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, ok)
}

func setupFile(t *testing.T, path string) {
	if err := os.Mkdir("/tmp/wo-fs/", os.ModePerm); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := os.RemoveAll("/tmp/wo-fs/"); err != nil {
			log.Println(err)
		}
	})

	if err := ioutil.WriteFile(path, []byte("test"), os.ModePerm); err != nil {
		t.Fatal(err)
	}
}

func setupDir(t *testing.T, path string) {
	if err := os.Mkdir("/tmp/wo-fs/", os.ModePerm); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := os.RemoveAll("/tmp/wo-fs/"); err != nil {
			log.Println(err)
		}
	})

	if err := os.Mkdir(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
