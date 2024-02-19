package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	message := greetings.Hello("khalil")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
