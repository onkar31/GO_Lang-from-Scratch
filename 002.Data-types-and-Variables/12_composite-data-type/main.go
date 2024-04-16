package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Now()
	year, month, day := t1.Date()
	fmt.Println(year, month, day)

}
