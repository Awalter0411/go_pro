package main

import "fmt"

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func compare(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func mutReturn() (name string, age int) {
	name = "tom"
	age = 20
	return name, age
}

func f3(args ...int) {

}

func main() {
	fmt.Printf("sum(1, 1): %v\n", sum(1, 1))
	name, age := mutReturn()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	f3(1, 2, 3)
}
