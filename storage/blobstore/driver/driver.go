package driver

type Driver interface {
	Read(path string) ([]byte, error)

	Write(path string, content []byte) error
}
