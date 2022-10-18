package main

import "fmt"

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
