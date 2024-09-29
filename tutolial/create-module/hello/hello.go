package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := os.Args[1:]

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
