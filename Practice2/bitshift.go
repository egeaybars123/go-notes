package main

import (
	"fmt"
	"strconv"
)

var x uint = 9  // 1001
var y uint = 33 // 100001
var z uint

func bitshift() {
	z = x << 1
	fmt.Println("x << 1", z)

	z = y >> 1
	fmt.Println("y >> 1", z)
}

func DecToBinary(number uint) string {
	output := ""
	for {
		if number/2 == 0 {
			output += "1"
			break
		}

		remainder := number % 2
		number = number / 2

		output += strconv.Itoa(int(remainder))

	}
	return reverseString(output)
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
