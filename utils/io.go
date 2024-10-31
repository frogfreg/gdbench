package utils

import (
	"fmt"
	"io"
	"os"
	"time"
)

const Mebibyte = 1024 * 1024
const Gibibyte = Mebibyte * 1024

const Megabyte = 1_000_000
const Gigabyte = Megabyte * 1_000

func SeqW(filename string) {
	samples := []float64{}

	f, err := os.OpenFile(filename, os.O_RDWR, 0o666)
	if err != nil {
		panic(err)
	}

	for range 5 {
		start := time.Now()
		n, err := f.WriteAt(make([]byte, Gigabyte), 0)
		finish := time.Since(start)
		if err != nil {
			panic(err)
		}

		speedSample := (float64(n) / float64(Megabyte)) / finish.Seconds()

		samples = append(samples, speedSample)
	}

	fmt.Printf("samples: %#v\n", samples)
}

func SeqR(filename string) {
	samples := []float64{}

	for range 5 {
		start := time.Now()

		data, err := os.ReadFile(filename)
		finish := time.Since(start)
		if err != nil && err != io.EOF {
			panic(err)
		}

		speedSample := (float64(len(data)) / float64(Megabyte)) / finish.Seconds()
		samples = append(samples, speedSample)
	}

	fmt.Printf("read samples: %#v\n", samples)
}
