package main

import (
	"fmt"
	"log"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)

	// A slice of names.
	names := []string{"(1)Mark", "(2)Aleaf", "(3)Charlie"}
	//var message string
	//var err error
	//var message, err = greetings.Hello("Mark")

	// Request greeting messages for the names.
	//var messages = []string
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(message)
	fmt.Println(messages)
	fmt.Println(quote.Go())
}
