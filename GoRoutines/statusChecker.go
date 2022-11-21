package main

import (
	"fmt"
	"net/http"
	"time"
)

// main routine
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://www.udemy.com",
		"http://www.github.com/B-Shiva-Kumar/",
	}

	// channel implementation
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// 1
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// 2
	// infinite loop
	for l := range c {
		// go checkLink(l, c)
		// anonmyous func / function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}

}

// go routine (child routine)
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down I think!"
		c <- link
		return
	} else {
		fmt.Println(link, "is up & running!")
		// c <- "Yep, its Up & running"
		c <- link
	}

}
