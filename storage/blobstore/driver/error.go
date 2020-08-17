package driver

import "errors"

var (
	ErrPathIsDir = errors.New("file to read is a directory")
	//ErrBlobAlreadyExists = errors.New("blob already exists")
)
