package files_test

import (
	"testing"

	"webtyp.com/files"
)

func TestErrNotExist_IsComparableAndDescribed(t *testing.T) {
	var err error = files.ErrNotExist
	if err != files.ErrNotExist {
		t.Fatal("ErrNotExist must compare equal to itself")
	}
	if err.Error() != "files: file does not exist" {
		t.Fatalf("message = %q", err.Error())
	}
}
