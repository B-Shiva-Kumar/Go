package main

import "fmt"

// pck lvl declaration
var (
	API    string = "https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy"
	myName string = "Shiva Kumar B"
)

// Variable Scope Visibility

// var starts with upper case
// visible to all the packages
var A string = "Global lvl"

// var starts with lower case
// visible to current package
var a string = "Package lvl"

func main() {

	fmt.Println("A - ", A)
	fmt.Println("a - ", a)

	// function lvl declaration
	// visible to the block or funtion lvl
	var a string = "block lvl"
	fmt.Println("a - ", a)

	var theURL string = "github.com/B-Shiva-Kumar/Go"

	fmt.Println(theURL, "\n")
	fmt.Println(API, "\n")
}
