package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3}
	var pa1 [3]*int
	for i := 0; i < len(a1); i++ {
		pa1[i] = &a1[i]
	}
	fmt.Printf("pa1: %v\n", pa1)
	for i := 0; i < len(pa1); i++ {
		fmt.Println(*pa1[i])
	}
}
