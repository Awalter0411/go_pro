package main

import "fmt"

type Person struct {
	name string
}

func show_person1(p Person) {
	p.name = "tom"
}

func show_person2(p *Person) {
	p.name = "tom"
}

func (p Person) show_person1() {
	p.name = "tom"
}

func (p *Person) show_person2() {
	p.name = "tom"
}

func main() {
	var p1 = Person{name: "curry"}
	var p2 = &Person{name: "james"}
	// fmt.Printf("p1: %T\n", p1)
	// fmt.Printf("p2: %T\n", p2)

	// show_person1(p1)
	// show_person2(p2)

	p1.show_person1()
	p2.show_person2()

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
}
