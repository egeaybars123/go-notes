package main

import "fmt"

const greeting = "Hello,"

var coin = "Ethereum"

func main() {
	fmt.Println(greeting, coin)
	setCoin("Bitcoin")
	fmt.Println(greeting, coin)
}

func setCoin(name string) {
	coin = name
}
