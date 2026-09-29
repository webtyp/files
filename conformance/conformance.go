// Package conformance is the test suite every implementation of the files contract runs
// against itself, so any two implementations are substitutable.
package conformance

import (
	"testing"

	"webtyp.com/files"
)

// Factory builds a fresh, empty implementation for one subtest.
type Factory struct {
	Name string
	New  func(t *testing.T) files.ReadWriter
}

// Run runs the ReadWriter suite and, when the implementation also is a files.Appender, the
// append suite.
func Run(t *testing.T, f Factory) {
	if f.New == nil {
		t.Fatal("conformance: Factory.New is required")
	}
	t.Run("read_missing_is_ErrNotExist", func(t *testing.T) {
		if _, err := f.New(t).ReadFile("missing.bin"); err != files.ErrNotExist {
			t.Fatalf("ReadFile(missing) error = %v, want files.ErrNotExist (unwrapped)", err)
		}
	})
	t.Run("write_then_read", func(t *testing.T) {
		rw := f.New(t)
		mustWrite(t, rw, "a.txt", "hello")
		expect(t, rw, "a.txt", "hello")
	})
	t.Run("write_replaces", func(t *testing.T) {
		rw := f.New(t)
		mustWrite(t, rw, "a.txt", "first, longer")
		mustWrite(t, rw, "a.txt", "second")
		expect(t, rw, "a.txt", "second")
	})
	t.Run("empty_file_exists", func(t *testing.T) {
		rw := f.New(t)
		mustWrite(t, rw, "empty", "")
		got, err := rw.ReadFile("empty")
		if err != nil || len(got) != 0 {
			t.Fatalf("ReadFile(empty) = %q, %v; want empty content and nil error", got, err)
		}
	})
	t.Run("paths_are_independent", func(t *testing.T) {
		rw := f.New(t)
		mustWrite(t, rw, "a", "A")
		mustWrite(t, rw, "b", "B")
		expect(t, rw, "a", "A")
		expect(t, rw, "b", "B")
	})
	t.Run("caller_buffer_is_not_shared", func(t *testing.T) {
		rw := f.New(t)
		data := []byte("abc")
		if err := rw.WriteFile("a", data); err != nil {
			t.Fatal(err)
		}
		data[0] = 'X'
		expect(t, rw, "a", "abc")
		got, _ := rw.ReadFile("a")
		got[0] = 'Y'
		expect(t, rw, "a", "abc")
	})
	if _, ok := f.New(t).(files.Appender); !ok {
		return
	}
	t.Run("append_creates", func(t *testing.T) {
		rw := f.New(t)
		if err := rw.(files.Appender).AppendFile("log", []byte("one")); err != nil {
			t.Fatal(err)
		}
		expect(t, rw, "log", "one")
	})
	t.Run("append_adds_to_end", func(t *testing.T) {
		rw := f.New(t)
		mustWrite(t, rw, "log", "one")
		if err := rw.(files.Appender).AppendFile("log", []byte(",two")); err != nil {
			t.Fatal(err)
		}
		expect(t, rw, "log", "one,two")
	})
}

func mustWrite(t *testing.T, w files.Writer, path, content string) {
	t.Helper()
	if err := w.WriteFile(path, []byte(content)); err != nil {
		t.Fatalf("WriteFile(%q): %v", path, err)
	}
}

func expect(t *testing.T, r files.Reader, path, want string) {
	t.Helper()
	got, err := r.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q): %v", path, err)
	}
	if string(got) != want {
		t.Fatalf("ReadFile(%q) = %q, want %q", path, got, want)
	}
}
