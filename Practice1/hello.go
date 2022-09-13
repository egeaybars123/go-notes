package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Ethereum")
	fmt.Println(message)
	fmt.Println(Greet("Litecoin"))
}
