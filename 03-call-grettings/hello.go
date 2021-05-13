package main

import (
	"fmt"
	"../02-greetings"
)

func main() {
	message := greetings.Hello("Omer")
	fmt.Println(message)
}