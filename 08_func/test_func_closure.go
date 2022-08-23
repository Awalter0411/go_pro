package main

import "fmt"

func add() func(int) int {
	x := 0
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	f := add()
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(20): %v\n", f(20))
	fmt.Printf("f(30): %v\n", f(30))
}
