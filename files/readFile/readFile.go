package main

import (
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(file)

	// 1
	// readFile, err := ioutil.ReadAll(file)
	// fmt.Print(string(readFile))

	// 2
	io.Copy(os.Stdout, file)

	// fmt.Println(os.Args)

}
