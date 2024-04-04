package main

import "fmt"

func main() {
	// shorthand for
	// var message string = "I love Go!"
	message := "I love Go!"

	// another way would be
	// var message string
	// message = "I love Go!"

	fmt.Println(message)

	// Variables are mean to store values which can change over time
	// Here, I'm changing the text which is stored in "message"
	message = "because it is a great programming language."

	// Although I'm executing the same instruction
	// the value stored in "message is different"
	fmt.Println(message)
}
