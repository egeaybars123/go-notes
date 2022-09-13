package greetings

import "fmt"

func Hello(coin string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", coin)
	return message
}
