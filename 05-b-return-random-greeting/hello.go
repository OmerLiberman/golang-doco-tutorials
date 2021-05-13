package main

import (
	"../05-a-return-random-greeting"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Greetings:")
	log.SetFlags(0)

	message, err := greetings.Hello("Omer")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}