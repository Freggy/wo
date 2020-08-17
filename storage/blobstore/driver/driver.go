package driver

// Driver implements the concrete logic on how to read/write blob data.
type Driver interface {

	// Read reads a data from the specified path.
	Read(path string) ([]byte, error)

	// Write writes data to the specified path.
	Write(path string, content []byte) error

	// Exists whether a given path exists.
	Exists(path string) (bool, error)
}
