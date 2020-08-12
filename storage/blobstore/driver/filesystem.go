package driver

import (
	"errors"
	"io/ioutil"
	"os"
)

var (
	ErrReadPathIsDir = errors.New("file to read is a directory")
	ErrBlobAlreadyExists = errors.New("blob already exists")
)

type filesystemDriver struct {
}

func (fd *filesystemDriver) Read(path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, ErrReadPathIsDir
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (fd *filesystemDriver) Write(path string, content []byte) error {
	if _, err := os.OpenFile(path, os.O_CREATE, os.ModePerm); os.IsExist(err) {
		return ErrBlobAlreadyExists
	}

	if err := ioutil.WriteFile(path, content, os.ModePerm); err != nil {
		return err
	}

	return nil
}