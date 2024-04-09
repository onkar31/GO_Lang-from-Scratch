package main

import "fmt"

//package level variable
var myName string 

func main() {

	var whatToSay string
	whatToSay = "Goodbye!"
	var i int

	fmt.Println(whatToSay)

	i = 8
	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()
	fmt.Println("The function  returned", whatWasSaid)

	firstThningThatWasSaid, theOtherThingThatWasSaid := saySomethingAgain()
	fmt.Println("THIS TIME THE FUNCTION RETURNED", firstThningThatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() string {
	return "something"
}

func saySomethingAgain() (string, string){
	return "SOMETHING", "ELSE!"
}
