package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for link := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(link)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}
	fmt.Println(link, "is up!")
	ch <- link
}
