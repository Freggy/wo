package driver

import (
	"io/ioutil"
	"os"
)

type filesystemDriver struct {
}

func (fd *filesystemDriver) Read(path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, ErrPathIsDir
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (fd *filesystemDriver) Write(path string, content []byte) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return ErrPathIsDir
	}

	if err := ioutil.WriteFile(path, content, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (fd *filesystemDriver) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) { // does not exist
		return false, nil
	} else if !os.IsNotExist(err) { // does exist
		return true, nil
	} else {
		return false, err
	}
}
