package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spahishBot struct{}

func main() {
	eb := englishBot{}
	sb := spahishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spahishBot) getGreeting() string {
	return "Hola!"
}
