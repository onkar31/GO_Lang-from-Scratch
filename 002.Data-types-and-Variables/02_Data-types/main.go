package main

import (
	"fmt"
	"time"
)

func main() {
// Basic Data types
	
	//string = "Go is Amazing"
	message := "Go is Amazing"
	fmt.Println(message)

	//whole numbers = 14
	var number int = 14
	fmt.Println(number)

	//real numbers = 1.43
	var realNumber float32 = 1.43
	fmt.Println(realNumber)

	//characters = s, !
	var ch rune = 's'
	fmt.Println(ch)

	//boolean = true/false 
	var isTrue bool = true 
	fmt.Println(isTrue)

	//composite data types
	var result time.Time = time.Now()
	fmt.Println(result)
	
}
