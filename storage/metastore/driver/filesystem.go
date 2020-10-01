package driver

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var ErrKeyNotFound = errors.New("key not found")

type filesystemDriver struct {
	basePath string
}

func (fd *filesystemDriver) Get(repository, mapName, key string) (interface{}, error) {
	path := fmt.Sprintf("%s/%s/%s", fd.basePath, repository, mapName)
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(f)
	for err != io.EOF {
		data, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}

		// lines have the following format:
		// key=value\n
		// key=value\n

		parts := strings.Split(data, "=")
		if len(parts) < 2 {
			continue
		}

		if parts[0] == key {
			return parts[1], nil
		}
	}
	return nil, ErrKeyNotFound
}

func (fd *filesystemDriver) Set(repository, mapName, key string, value interface{}) error {
	f, err := fd.openFile(repository, mapName, key)
	if err != nil {
		return err
	}
	defer f.Close()
	kv := make(map[string]interface{})
	r := bufio.NewReader(f)
	// fill key-value pair map
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		parts := strings.Split(line, "=")
		if len(parts) < 2 {
			continue
		}
		kv[parts[0]] = parts[1]
	}
	kv[key] = value
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("truncate file: %s", err)
	}
	w := bufio.NewWriter(f)
	for k, v := range kv {
		if _, err := w.WriteString(fmt.Sprintf("%s=%s\n", k, v)); err != nil {
			return fmt.Errorf("write string: %s", err)
		}
	}
	if err := w.Flush(); err != nil {
		return fmt.Errorf("flush file: %s", err)
	}
	return nil
}

func (fd *filesystemDriver) openFile(repository, mapName, key string) (*os.File, error) {
	path := fmt.Sprintf("%s/%s/%s", fd.basePath, repository, mapName)
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
