package driver

import "errors"

var (
	ErrPathIsDir = errors.New("file to read is a directory")
)
