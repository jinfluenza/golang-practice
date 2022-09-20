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
		"https://amazon.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, site := range links {
		go makeRequest(site, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			makeRequest(link, c)
		}(l)
	}
}

func makeRequest(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Printf("%s is down! \n", site)
		c <- site
	} else {
		fmt.Printf("%s is up! \n", site)
		c <- site
	}
}
