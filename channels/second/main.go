package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) { //pass copy of l variable in order not to have the same variable (l) in different go routines
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //function literal (lambda)
		//fix in go 1.22
		//https://go.dev/blog/go1.22
		//https://go.dev/blog/loopvar-preview
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
