package utils

import (
	"io"
	"os"
)

func SequentialWrite(dir string, size int) (string, error) {
	f, err := os.CreateTemp(dir, "tempFile-seq")
	if err != nil {
		return "", err
	}
	data := make([]byte, size)

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}

	return f.Name(), f.Close()
}

func SequentialRead(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 4096)

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
	}

	return nil
}
