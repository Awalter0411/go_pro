package main

import "fmt"

func main() {
	var b1 bool = true
	b2 := false
	age := 18

	if age > 18 {
		fmt.Print("你已经成年了")
	} else {
		fmt.Print("未成年")
	}
}
