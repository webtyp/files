package mem_test

import (
	"testing"

	"webtyp.com/files"
	"webtyp.com/files/conformance"
	"webtyp.com/files/mem"
)

func TestMemConformance(t *testing.T) {
	conformance.Run(t, conformance.Factory{
		Name: "mem",
		New:  func(t *testing.T) files.ReadWriter { return mem.New() },
	})
}
