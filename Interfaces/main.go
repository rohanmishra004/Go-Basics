package main

import "fmt"

//we declare the bot interface - new custom bot
type bot interface {
	getGreeting() string
}

//we have created two bots
type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

//getGreeting function for both spanish and english bot is different
func (englishBot) getGreeting() string {
	//custom logic for english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	//custom logic for spanish greeting
	return "Hola!"
}

//but here for printGreeting - we would use interface because the logic for both these functions is same

//here we would like to compress both these functions together

//we will now refactor both printGreeting functions
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
