package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://asfdaflsaf;jsa.com",
	}

	for _, site := range links {
		makeRequest(site)
	}
}

func makeRequest(site string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Printf("%s is down! \n", site)
	} else {
		fmt.Printf("%s is up! \n", site)
	}
}
