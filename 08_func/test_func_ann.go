package main

import "fmt"

func main() {
	r := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 2)
	fmt.Printf("r: %v\n", r)

	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	fmt.Printf("max(1, 2): %v\n", max(1, 2))
}
