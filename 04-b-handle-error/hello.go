package main

import (
	"fmt"
	"log"
	"../04-a-return-handle-error"
)

func main() {
	log.SetPrefix("greetings: ")
	//log.SetFlags(0) // flag to disable printing the time, source file, and line number.

	msg, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}