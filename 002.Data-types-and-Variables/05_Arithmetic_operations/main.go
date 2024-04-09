package main

import "fmt"

func main() {
	num1 := 15
	num2 := 60

	sum := num1 + num2
	subtraction := num1 - num2
	multiplication := num1 * num2
	division := num2 / num1
	modulus := num2 % num1
	fmt.Println("Sum:", sum, "\nSubtraction:", subtraction, "\nMultiplication:", multiplication,
		"\nDivision:", division, "\nRemainder:", modulus)

	//these operations change the values stored inside the variables
	num1++ // 15 + 1
	num2-- // 60 - 1
	fmt.Println(num1, num2)

}
