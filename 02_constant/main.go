package main

import (
	"fmt"
)

func main() {
	// const name string = "tom"
	// const age int = 20

	// const (
	// 	name = "tom"
	// 	age  = 20
	// )

	// const w, h = 200, 300

	// iota
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1) // 0
	fmt.Printf("a1: %v\n", a2) // 1
	fmt.Printf("a1: %v\n", a3) // 2

	const (
		b1 = iota
		_
		b2 = iota
	)
	fmt.Printf("b1: %v\n", b1) // 0
	fmt.Printf("b2: %v\n", b2) // 2

	const (
		c1 = iota
		c2 = 100
		c3 = iota
	)
	fmt.Printf("c1: %v\n", c1) // 0
	fmt.Printf("c1: %v\n", c2) // 100
	fmt.Printf("c1: %v\n", c3) // 2
}
