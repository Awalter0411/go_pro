package main

import "fmt"

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	street string
	city   string
}

func main() {
	var address = Address{
		street: "gsw",
		city:   "usa",
	}
	var tom = Person{
		name:    "curry",
		age:     34,
		address: address,
	}
	fmt.Printf("tom: %v\n", tom)
}
