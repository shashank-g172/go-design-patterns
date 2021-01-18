package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	c := make(chan string)
	for _, link := range links {
		// Create new subroutine for every GET request
		go checkLink(link, c)
	}

	// Inifite loop - watch for channel c, assign to l
	for l := range c {
		// call the subroutine again using a function literal. Do not try to access a variable from the main routine,
		// always use the variables returned from the child routine using the channel
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// Blocking call
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}
