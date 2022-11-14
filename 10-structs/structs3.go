package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// sigle struct
type Contact struct {
	id    int
	name  string
	phone int
}

// nested struct
type ShowRoom struct {
	address    string
	location   string
	postalcode int
}

// sub nested struct
type Car struct {
	id       int
	name     string
	price    float32
	showroom ShowRoom
}

// same struct type defination
type Person struct {
	Firstname, Lastname string
}

// json struct
type Book struct {
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Author string  `json:"author"`
}

// struct as parameter
func GetContact(contact Contact) {
	fmt.Println("contact details: ", contact)
}

// read struct with method
func (supercar Car) GetSuperCar() {
	fmt.Println("read struct with method", supercar)
}

func main() {

	var ContactPerson Contact
	var SuperCar Car
	var Biodata Person

	// abstractStruct := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "john doe",
	// 	Age:  23,
	// }
	var MyBook Book

	ContactPerson.id = 1
	ContactPerson.name = "john doe"
	ContactPerson.phone = 87887242895
	NewContactPerson := Contact{2, "jane doe", 8777888555666}
	// GetContactPointer := &ContactPerson

	SuperCar.id = 1
	SuperCar.name = "lambogini"
	SuperCar.price = 1.500000000
	SuperCar.showroom.address = "jl.nusantara kap 50 - 51"
	SuperCar.showroom.location = "depok"
	SuperCar.showroom.postalcode = 163353

	Biodata.Firstname = "aldi"
	Biodata.Lastname = "khan"

	MyBook.Name = "hary potter"
	MyBook.Price = 50.000
	MyBook.Author = "jk.rolling"

	json, err := json.Marshal(MyBook)

	fmt.Println("read struct step one", ContactPerson)
	fmt.Println("read struct step two", NewContactPerson)
	// fmt.Printf("read struct by key %v\n", ContactPerson.name)
	GetContact(ContactPerson)
	// fmt.Println("read struct with pointer", *GetContactPointer)
	// fmt.Println("read struct by key with pointer", (*GetContactPointer).phone)
	// fmt.Println("read nested struct", SuperCar)
	// fmt.Println("read nested struct with key", SuperCar.showroom.location)
	SuperCar.GetSuperCar()
	// fmt.Println("read struct with same type defination", Biodata)
	// fmt.Println("read abstract struct", abstractStruct)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("read struct with json style", string(json))
	}
}
