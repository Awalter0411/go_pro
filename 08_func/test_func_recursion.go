package main

import "fmt"

func f1(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * f1(n-1)
	}
}

func f2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return f2(n-1) + f2(n-2)
}

func main() {
	fmt.Printf("f1(5): %v\n", f1(5))
	fmt.Printf("f2(5): %v\n", f2(5))
}
