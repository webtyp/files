// Package files is the contract for reading and writing whole files by path. A consumer asks
// for the narrowest interface it needs (a model loader asks for a Reader and can never write);
// the composition root passes an implementation: webtyp/opfs in a browser Worker, a disk
// implementation on a server, files/mem in tests.
package files

// Reader reads a whole file.
type Reader interface {
	// ReadFile returns the contents of the file at path. When path does not exist it returns
	// exactly ErrNotExist (the value itself, never wrapped), so callers compare with ==.
	ReadFile(path string) ([]byte, error)
}

// Writer writes a whole file.
type Writer interface {
	// WriteFile replaces the contents of the file at path with data, creating it if needed.
	// The implementation keeps its own copy: changing data afterwards changes nothing stored.
	WriteFile(path string, data []byte) error
}

// Appender adds bytes to the end of a file.
type Appender interface {
	// AppendFile adds data to the end of the file at path, creating it if needed.
	AppendFile(path string, data []byte) error
}

// Remover deletes a file.
type Remover interface {
	// RemoveFile deletes the file at path. When path does not exist it returns exactly
	// ErrNotExist, like ReadFile, so a caller that only wants the file gone compares with ==.
	RemoveFile(path string) error
}

// ReadWriter reads and writes whole files.
type ReadWriter interface {
	Reader
	Writer
}

// ErrNotExist is what ReadFile returns when the file does not exist.
var ErrNotExist error = notExist{}

type notExist struct{}

func (notExist) Error() string { return "files: file does not exist" }
