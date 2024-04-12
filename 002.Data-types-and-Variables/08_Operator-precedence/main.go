package main

import "fmt"

func main() {
	x := 5
	y := 2
	z := 3

	result := x + y * z //Multiplication has higher precedence than Addition 
	fmt.Println(result) // Output: 11

	result = (x + y) * z //Parantheses can be used to explicitly define precedence 
	fmt.Println(result) // Output: 21

	result = x * y / z //Multiplication and division have equal precedence and are evaluated from left to right
	fmt.Println(result) // Output: 3

	result = (x * y) / z //Parantheses can also be used to specify evaluation order
	fmt.Println(result) // Output: 3

	// More examples 
	a := true 
	b := false 
	c := true

	resultBool := a || b && c // && has higher precedence than ||
	fmt.Println(resultBool)   // Output: true
}