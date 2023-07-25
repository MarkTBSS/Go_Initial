package main

import (
	"fmt"

	"example/greetings"

	"rsc.io/quote"
)

func main() {
	//var message string
	message := greetings.Hello("Mark")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
