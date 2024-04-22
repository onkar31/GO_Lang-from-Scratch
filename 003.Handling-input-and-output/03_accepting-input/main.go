package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
}