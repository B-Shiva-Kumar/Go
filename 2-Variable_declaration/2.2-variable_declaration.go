package main

import "fmt"

// Variables declaration can be grouped together into blocks for greater readability and code quality.

// pck lvl variable declaration

// multiple variables declaration at one time
var (
	id   int    = 1
	name string = "Go gopher"
	lang string = "golang"
)

var (
	count int = 1
)

func main() {

	// fmt.Println(count)
	// fmt.Println(id, name, lang)
	fmt.Println(id, name, lang, count)

}
