// Package mem is the in-memory reference implementation of the files contract, for tests and
// demos. It keeps nothing after the process ends.
package mem

import (
	"sync"

	"webtyp.com/files"
)

type entry struct {
	path string
	data []byte
}

// Files keeps whole files in memory. The zero value is not usable; call New.
type Files struct {
	mu      sync.Mutex
	entries []entry
}

// New returns an empty in-memory file set.
func New() *Files { return &Files{} }

var (
	_ files.ReadWriter = (*Files)(nil)
	_ files.Appender   = (*Files)(nil)
	_ files.Remover    = (*Files)(nil)
)

func (f *Files) find(path string) int {
	for i := range f.entries {
		if f.entries[i].path == path {
			return i
		}
	}
	return -1
}

// ReadFile returns a copy of the file, or files.ErrNotExist.
func (f *Files) ReadFile(path string) ([]byte, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	i := f.find(path)
	if i < 0 {
		return nil, files.ErrNotExist
	}
	return append([]byte(nil), f.entries[i].data...), nil
}

// WriteFile stores a copy of data at path.
func (f *Files) WriteFile(path string, data []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	cp := append([]byte(nil), data...)
	if i := f.find(path); i >= 0 {
		f.entries[i].data = cp
		return nil
	}
	f.entries = append(f.entries, entry{path: path, data: cp})
	return nil
}

// AppendFile adds a copy of data to the end of the file at path.
func (f *Files) AppendFile(path string, data []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if i := f.find(path); i >= 0 {
		f.entries[i].data = append(f.entries[i].data, data...)
		return nil
	}
	f.entries = append(f.entries, entry{path: path, data: append([]byte(nil), data...)})
	return nil
}

// RemoveFile deletes the file at path, or returns files.ErrNotExist.
func (f *Files) RemoveFile(path string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	i := f.find(path)
	if i < 0 {
		return files.ErrNotExist
	}
	f.entries = append(f.entries[:i], f.entries[i+1:]...)
	return nil
}
