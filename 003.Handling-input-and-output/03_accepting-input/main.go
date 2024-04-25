package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	sum := num1 + num2
	fmt.Println("Addition of two numbers", sum)

	//-----------------------------------------------------//

	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSpace(line)
	fmt.Println("Line 1:", line)

	// Enter a number
	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSpace(line)

	//convert string into int
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	fmt.Println("My number is", num)
	fmt.Println("Increment it by 10:", num+10)

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	floatNum, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("My float number is:", floatNum)
}
