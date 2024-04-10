package main

import "fmt"

func main() {

	var smallInt int8 = 100
	smallInt = smallInt + 28
	fmt.Println("Example of Overflow", smallInt)

	var verySmallInt int8 = -127
	verySmallInt = verySmallInt - 2
	fmt.Println("Example of Underflow", verySmallInt)

	var toMultiply int8 = 60
	toMultiply = toMultiply * 3
	fmt.Println("Example of Overflow via Multiplication", toMultiply)

	var toDivide int8 = 15
	toDivide = toDivide / 2
	fmt.Println("Example of whole number Division", toDivide)

	var toMod int8 = 15
	toMod = toMod % 10
	fmt.Println("Example of Modulus operation", toMod)

	var toModNegative int8 = -15
	toModNegative = toModNegative % 2
	fmt.Println("Example of Modulus operation with Negative value", toModNegative)

	var divideByZero int = 10
	//divideByZero = divideByZero / 0		// will give compile time error
	fmt.Println(divideByZero)

}
