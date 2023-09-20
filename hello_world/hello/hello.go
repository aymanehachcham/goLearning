package main

import (
	"examples/custom_hello"
	"fmt"
	"log"
)

func main() {

	// Setting logging levels
	log.SetPrefix("custom_hello: ")
	log.SetFlags(0)

	names := []string{"Ayamne", "Lola", "Zobi"}
	messages, err := custom_hello.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
