package main

import "fmt"

type Person struct {
	name string
	age  int
}

func show_person(p Person) {
	p.name = "curry"
	fmt.Printf("p.name: %v\n", p.name)
}

func show_person2(p *Person) {
	p.name = "james"
	fmt.Printf("p.name: %v\n", p.name)
}

func main() {
	var p = Person{
		name: "tom",
		age:  34,
	}
	fmt.Printf("p.name: %v\n", p.name)
	// 值传递
	show_person(p)
	fmt.Printf("p.name: %v\n", p.name)

	// 指针传递
	show_person2(&p)
	fmt.Printf("p.name: %v\n", p.name)
}
