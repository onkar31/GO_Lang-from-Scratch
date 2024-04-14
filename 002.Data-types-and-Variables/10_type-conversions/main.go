package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int8 = 42
	var biggerNum int16 = int16(num)
	fmt.Println("Widening:", biggerNum)

	//Data loss -> Narrow Casting
	var bigNum int16 = 300
	var smallNum int8 = int8(bigNum)
	fmt.Println("Narrow Casting:", smallNum)

	// Integers can be safely converted to real numbers without losing any data
	var anInt int16 = 31
	var aFloat float32 = float32(anInt)
	fmt.Println("Integer to Float Type Conversion:", aFloat)

	// Float to Integer Conversion
	aFloat = 3.5
	var anotherInt int = int(aFloat)
	fmt.Println("Float to Integer Type Conversion:", anotherInt)

	// string to integer conversion
	//var numAsString string = "22"
	num1, err := strconv.Atoi("22")
	fmt.Println("String to Integer Type Conversion", num1, err)

	// Integer to string conversion
	var numAsString string = strconv.Itoa(num1)
	fmt.Println("Integer to String Type Conversion:", numAsString)
}
