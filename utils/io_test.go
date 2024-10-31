package utils

import (
	"os"
	"testing"
)

func TestSeqW(t *testing.T) {

	f, err := os.CreateTemp("testfiles", "test-seqw-*")
	if err != nil {
		t.Error(err)
	}

	SeqW(f.Name())

	if err := os.Remove(f.Name()); err != nil {
		t.Error(err)
	}

	t.Error("erring on purpose")
}

func TestSeqR(t *testing.T) {

	f, err := os.CreateTemp("testfiles", "test-seqr-*")
	if err != nil {
		t.Error(err)
	}

	if _, err := f.Write(make([]byte, Gigabyte)); err != nil {
		t.Error(err)
	}

	if err := f.Close(); err != nil {
		t.Error(err)
	}

	SeqR(f.Name())

	if err := os.Remove(f.Name()); err != nil {
		t.Error(err)
	}

	t.Error("erring on purpose")
}
