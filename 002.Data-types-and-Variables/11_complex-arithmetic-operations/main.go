package main

import (
	"fmt"
	"math"
)

func main() {

	// Square root of a number
	var num int = 22
	result := math.Sqrt(float64(num))
	fmt.Println(result)

	// Absolute value of a number
	num = -14
	fmt.Println(math.Abs(float64(num)))

	// Exponential operations
	num = 9
	fmt.Println(math.Pow(float64(num), 3))
}
