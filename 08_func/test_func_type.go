package main

import "fmt"

type f1 func(int, int) int

func sum(a int, b int) int {
	return a + b
}

func main() {
	var f f1
	f = sum
	fmt.Printf("f(1, 2): %v\n", f(1, 2))
}
