package main

import "fmt"

type person struct {
	firstName, lastName string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	// embedding structs
	jim := person{
		firstName: "Jim",
		lastName:  "Hales",
		contactInfo: contactInfo{
			email:   "Jim@Hales.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("Jimmy")
	jim.print()

}

// update persons name
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// reciever func for person details
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
