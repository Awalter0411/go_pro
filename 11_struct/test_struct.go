package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Customer struct {
	id, age       int
	name, address string
}

func main() {
	var tom Person
	tom = Person{}
	tom.name = "curry"
	tom.age = 34
	fmt.Printf("tom: %v\n", tom)

	// 匿名结构体
	var dog struct {
		id   int
		name string
	}
	dog.name = "james"
	fmt.Printf("dog.name: %v\n", dog.name)
}
