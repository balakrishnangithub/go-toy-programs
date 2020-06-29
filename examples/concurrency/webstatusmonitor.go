package main

import (
	"fmt"
	"net/http"
	"time"
)

// WebStatus ...
type WebStatus struct {
	link   string
	isDown bool
}

func checkLink(link string, c chan WebStatus) {
	_, err := http.Get(link)
	if err != nil {
		c <- WebStatus{link: link, isDown: true}
	}
	c <- WebStatus{link: link, isDown: false}
}

func mainWebStatusMonitor() {
	links := []string{
		"https://golang.com",
		"https://tour.golang.org/",
		"https://play.golang.org/",
		"https://go.dev/",
		"https://pkg.go.dev/",
	}
	queue := make(chan WebStatus, len(links))
	for _, link := range links {
		go checkLink(link, queue) // Non-Blocking Call
	}
	for webStatus := range queue {
		fmt.Printf("%+v\n", webStatus)

		// time.Sleep is a Blocking Call, so its required to run it in a Goroutine.
		go (func(link string) {
			time.Sleep(3 * time.Second) // Blocking Call
			checkLink(link, queue)      // Blocking Call
		})(webStatus.link) // Non-Blocking Call
	}
	/*
		for i := 0; i < len(links); i++ {
			status := <-queue // Blocking Call
			fmt.Printf("%+v\n", status)
		}
	*/
}
