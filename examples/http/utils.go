package main

import (
	"bytes"
	"io"
	"os"
)

// coutAll reads from r and writes the value to stdout.
// Note: ioutil.ReadAll reads from r and returns the value.
func coutAll(r io.Reader) error {
	buf := make([]byte, bytes.MinRead) // bytes.MinRead is 512
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			os.Stdout.Write(buf[:n])
			break
		}
		if err != nil {
			return err
		}
		os.Stdout.Write(buf)
	}
	/*
		Alternatively we can use `bytes.Buffer`
		it automatically grows the buffer by `bytes.MinRead` until `io.EOF`

		var buf bytes.Buffer
		_, err := buf.ReadFrom(r)
		if err != nil {
			return err
		}
		os.Stdout.Write(buf.Bytes())
	*/
	return nil
}
