package utils

import (
	"bytes"
	"io"
	"os"
)

// CoutAll reads from r and writes the value to stdout.
// It returns the number of bytes written and an error, if any.
// Note: `ioutil.ReadAll` reads from r and returns the value.
func CoutAll(r io.Reader) (coutCount int, err error) {
	buf := make([]byte, bytes.MinRead) // bytes.MinRead is 512
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			wc, err := os.Stdout.Write(buf[:n])
			if err != nil {
				return coutCount, err
			}
			coutCount += wc
			break
		}
		if err != nil {
			return coutCount, err
		}
		wc, err := os.Stdout.Write(buf)
		if err != nil {
			return coutCount, err
		}
		coutCount += wc
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
	return coutCount, err
}
