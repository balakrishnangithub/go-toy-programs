package main

import (
	"log"
	"net/http"
)

func main() {
	mainIPGeolocationAPI()

	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	coutAll(resp.Body)
}
