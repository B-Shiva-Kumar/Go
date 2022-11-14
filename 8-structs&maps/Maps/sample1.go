package main

import "fmt"

func main() {
	m := map[string]string{
		"dog": "bark",
	}

	addToMap(m)

	fmt.Println(m)

	for _, value := range m {
		fmt.Println(value)
	}
}

func addToMap(m map[string]string) {
	m["cat"] = "purr"
}
