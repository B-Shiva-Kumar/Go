// simple get request

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// 1
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// 2. alternative
	// func Copy(dst Writer, src Reader) (written int64, err error)
	io.Copy(os.Stdout, resp.Body)
}
