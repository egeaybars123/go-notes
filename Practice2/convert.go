package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var strVar = "100"

func main() {
	//converts the string to int
	intVar, err := strconv.Atoi(strVar)

	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	//bitwise shifting operation
	bitshift()

	//Decimal to Binary converter
	fmt.Println(DecToBinary(68))
}
