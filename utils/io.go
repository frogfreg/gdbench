package utils

import (
	"fmt"
	"io"
	"os"
	"time"
)

const gigabyte = 1024 * 1024 * 1024

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

func BenchMarkSequentialWrite() {
	samples := []int64{}

	finishTime := time.Now().Add(10 * time.Second)

	f, err := os.CreateTemp("testfiles", "temp")
	if err != nil {
		panic(err)
	}

	writtenBytes := 0

	for time.Now().Before(finishTime) {

		if writtenBytes >= gigabyte {
			writtenBytes = 0
		}
		start := time.Now()
		n, err := f.WriteAt(make([]byte, 1024), int64(writtenBytes))
		if err != nil {
			panic(err)
		}
		writtenBytes += n
		finishTime := time.Since(start)

		samples = append(samples, int64(finishTime))
	}

	fmt.Printf("the length of samples is: %v\n", len(samples))
}
