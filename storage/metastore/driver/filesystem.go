package driver

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	ErrKeyNotFound = "key not found"
)

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

	// TODO:
	// * loop through lines
	// * if startsWith(key)
	// * split("=") -> return parts[1]

	return nil, nil
}

func (fd *filesystemDriver) Set(repository, mapName, key string, value interface{}) error {
	return nil
}
