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

	// what makes a valid variable name?
	// characters from a-z (capital case as well)
	// numbers (as long as it's not in the beginning)
	// underscores

	// Valid variable name: num, num1, n1n2n3, _n1, n1_and_n2
	// Invalid variable name: 1num, $num, @&!res

	// camelCase  -> firstName, lastName
	// PascalCase -> FirstName, LastName
	// snake_case -> first_name, last_name
	// UPPER_CASE -> FIRST_NAME, LAST_NAME
}
