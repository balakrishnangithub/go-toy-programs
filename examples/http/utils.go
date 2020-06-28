package main

import (
	"io"
	"os"
)

// coutAll reads from r and writes the value to stdout.
// Note: ioutil.ReadAll reads from r and returns the value.
func coutAll(r io.Reader) error {
	bufSize := 512
	buf := make([]byte, bufSize)
	for {
		n, err := r.Read(buf)
		if err != nil {
			return err
		}
		if n != bufSize {
			os.Stdout.Write(buf[:n])
			break
		}
		os.Stdout.Write(buf)
	}
	return nil
}
