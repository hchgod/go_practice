package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

func copySourceToDest(source io.Reader, dest io.Writer) error {
	return nil
}
func TestCopySourceToDest(t *testing.T) {
	const input = "foo"
	source := strings.NewReader(input)
	dest := bytes.NewBuffer(make([]byte, 0))

	err := copySourceToDest(source, dest)
	if err != nil {
		t.FailNow()
	}

	got := dest.String()
	if got != input {
		t.Errorf("expected: %s, got %s", input, got)
	}
}
