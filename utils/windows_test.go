package utils

import (
	"testing"
)

func TestGetVolumes(t *testing.T) {
	_, err := GetVolumes()
	if err != nil {
		t.Error(err)
	}
}
