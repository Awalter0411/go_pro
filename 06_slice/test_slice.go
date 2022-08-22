package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3}
	var s2 = make([]int, 2)
	fmt.Printf("a1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
}
