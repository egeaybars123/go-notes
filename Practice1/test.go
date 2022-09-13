package main

import "fmt"

func Greet(name string) string {
	message := fmt.Sprintf("I love %v ", name)
	return message
}
