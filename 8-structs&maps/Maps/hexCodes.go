package main

import "fmt"

func main() {
	// 1. empty map
	// var colors map[string]string
	// map[]

	// 2. declaring map
	var hexCodes = map[string]string{
		"red":   "#FF0000",
		"green": " #00FF00",
		"blue":  "0000FF",
	}
	fmt.Println(hexCodes)
	// map[blue:0000FF green: #00FF00 red:#FF0000]

	// 3. creating map using make func
	// we can only access in sqr braces not with . functionality
	var colors = make(map[string]string)
	colors["r"] = "red"
	colors["g"] = "green"
	colors["b"] = "blue"

	fmt.Println(colors)
	// map[b:blue g:green r:red]

	// delete key-value in maps
	delete(colors, "b")
	fmt.Println(colors)
	// map[b:blue g:green r:red]

}
