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
	var message string
	var err error
	message, err = greetings.Hello("Mark")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(quote.Go())
}
