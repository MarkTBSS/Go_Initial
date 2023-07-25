package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	var message string
	message = greetings.Hello("Gladys")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
