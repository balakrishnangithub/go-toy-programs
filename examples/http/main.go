package main

import (
	"log"
	"net/http"

	myutils "github.com/bkmagnetron/go-toy-programs/pkg/utils"
)

func main() {
	mainIPGeolocationAPI()

	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	myutils.CoutAll(resp.Body)
}
