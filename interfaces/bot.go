package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
	// getBotVersion()
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb) // Hi here
	printGreeting(sb) // Hola!

}
