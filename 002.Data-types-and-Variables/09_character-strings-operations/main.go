package main

import "fmt"

func main() {

	var ch rune = '5'
	fmt.Println("Value of ch:", ch)

	// Character have the same properties as whole numbers because they are whole number in memory
	// All oeprations on characters work the same way they do on whole numbers
	ch = ch + 1
	fmt.Println(ch)
	fmt.Println(ch / 2)
	fmt.Println(ch * 2)
	fmt.Println(ch % 2)
	fmt.Println("+++++++++++++++string operations+++++++++++++++++")

	// The only operation for string available is called "concatenation"
	// It allows us to "glue" two strings together
	var message string = "Go is "
	message = message + "awesome!" //concatenation
	fmt.Println(message)

	fmt.Println("++++++++++++++++Escaping Characters++++++++++++++++++")

	fmt.Println("Go is an \t awesome\\") // \t used for space
	// To print double quotes in print statment
	fmt.Println("The Person Who Masters Himself Through \"Self-Control\" and \"Discipline\" is Truly Undefeatable!")

}
