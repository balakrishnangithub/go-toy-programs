package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/dustin/go-humanize"
)

// WriteCounter counts the number of bytes written to it.
type WriteCounter struct {
	Total uint64 // Total # of bytes transferred
}

// Write implements the io.Writer interface.
//
// Always completes and never returns an error.
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	fmt.Printf("\r%v bytes", humanize.Bytes(wc.Total))
	return n, nil
}

func DownloadFile(url string) (err error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := path.Base(req.URL.Path)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, io.TeeReader(resp.Body, &WriteCounter{}))
	if err != nil {
		return err
	}
	fmt.Println()

	return
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(errors.New("URL argument is required and accepts only one URL."))
	}
	url := os.Args[1]

	err := DownloadFile(url)
	if err != nil {
		log.Fatal(err)
	}
}
