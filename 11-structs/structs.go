package main

import "fmt"

// defining Struct
type person struct {
	firstName, lastName string
}

func main() {
	// declare Struct
	alex := person{"Alex", "Anderson"}
	sam := person{firstName: "Sam", lastName: "Sundersan"}
	// fmt.Println(alex, sam)

	// {} or empty struct -> zer0 value for struct
	var Raju person
	fmt.Println(alex, sam, Raju)
	fmt.Printf("%+v\n", alex)
	fmt.Printf("%v\n", alex)

	// update value
	alex.firstName = "Asus"
	alex.lastName = "Alex"
	fmt.Println(alex)

}
