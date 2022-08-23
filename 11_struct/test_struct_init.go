package main

import "fmt"

type Person struct {
	id   int
	age  int
	name string
}

func main() {
	var tom Person
	tom = Person{
		1,
		18,
		"tom",
	}
	var curry = Person{
		name: "curry",
		id:   2,
		age:  34,
	}
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("curry: %v\n", curry)
}
