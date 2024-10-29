package utils

import (
	"os"
	"testing"
)

func TestSequentialWrite(t *testing.T) {
	tempDir, err := os.MkdirTemp("testfiles", "tempDir-seq")
	if err != nil {
		t.Error(err)
	}

	if _, err := SequentialWrite(tempDir, 1024*1024*100); err != nil {
		t.Error(err)
	}

	if err := os.RemoveAll(tempDir); err != nil {
		t.Error(err)
	}
}

func TestSequentialRead(t *testing.T) {
	tempDir, err := os.MkdirTemp("testfiles", "tempDir-seq")
	if err != nil {
		t.Error(err)
	}

	fName, err := SequentialWrite(tempDir, 1024*1024*100)
	if err != nil {
		t.Error(err)
	}

	if err := SequentialRead(fName); err != nil {
		t.Error(err)
	}

	if err := os.RemoveAll(tempDir); err != nil {
		t.Error(err)
	}
}

func TestBenchmarkSequentialWrite(t *testing.T) {
	BenchMarkSequentialWrite()
}
