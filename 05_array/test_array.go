package main

import "fmt"

func main() {
	var a1 [2]int
	var a2 = [2]int{1, 2}
	var a3 = [...]int{1, 2, 3, 4}
	var a4 = [...]int{1: 100, 2: 1200}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
}
