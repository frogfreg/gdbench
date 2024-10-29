package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetVolumes(t *testing.T) {
	volumes, err := GetVolumes()
	if err != nil {
		t.Error(err)
	}
	t.Log(volumes)

	fp, err := filepath.Abs(volumes[1])
	if err != nil {
		t.Error(err)
	}

	if err := os.WriteFile(filepath.Join(fp, "test.txt"), []byte("Hello there!!!"), 0o666); err != nil {
		t.Error(err)
	}

	t.Error("erring on purpose")
}
