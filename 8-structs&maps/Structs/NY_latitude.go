package main

import "fmt"

type location struct {
	longitude float64
	latitude  float64
}

func main() {
	newYork := location{
		latitude:  40.73,
		longitude: -73.93,
	}

	newYork.changeLatitude()

	fmt.Println(newYork)
}

func (lo *location) changeLatitude() {
	(*lo).latitude = 41.0
}

// The changeLatitude function expects a receiver of type pointer to a location struct ,
// but in the main function the receiver is a value type of a struct.
// What will happen when this code is executed?

// this program uses a shortcut where Go will automatically assume that even though
// newYork.changeLatitude() is using a value type we probably meant to pass in a pointer to the newYork struct.
