package main

import "fmt"

func main() {
	var aFloat float32 = 14.2

	fmt.Println("Addition:", aFloat+3)
	fmt.Println("Substraction:", aFloat-2)
	fmt.Println("Multiplication:", aFloat*3)
	fmt.Println("Division", aFloat/3)

	// modulus is not defined for floats
	//fmt.Println(aFloat % 3)

	var bFloat float32 = aFloat
	aFloat++
	bFloat--
	fmt.Println("After increament:", aFloat)
	fmt.Println("After Decreament:", bFloat)

	//divide by zero
	fmt.Println("++++++++++++++Divide by Zero+++++++++++++++")
	var toDivide = 3.0
	fmt.Println(toDivide / 0.0)
	fmt.Println(toDivide / -0.0)
	fmt.Println(-toDivide / 0.0)
	fmt.Println(-toDivide / -0.0)
}
