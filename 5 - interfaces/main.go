package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	//sb := spanishBot{}

	printGreeting((eb))

}
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
