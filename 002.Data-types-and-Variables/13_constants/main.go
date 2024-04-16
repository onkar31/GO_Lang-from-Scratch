package main

import (
	"fmt"
	"time"
)

const (
	// Declaring multiple constants
	companyName = "Acme Corporation"
	yearFounded = 2000
)

func main() {
	// Using the constants
	fmt.Println("Company:", companyName)
	fmt.Println("Year Founded:", yearFounded)

	// This will result in a compilation error since constants cannot be reassigned.
	// companyName = "New Company"
	
	// This will result in a compilation error since constants cannot accepts/assigned with composite data type.
	//const time = time.Now()
}
