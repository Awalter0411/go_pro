package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("hello %v\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func calc(calc_type string) func(int, int) int {
	switch calc_type {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	test("curry", sayHello)
	fmt.Printf("calc(\"+\")(1, 1): %v\n", calc("+")(1, 1))
}
