package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p_person *Person
	var tom = Person{
		name: "tom",
		age:  30,
	}
	p_person = &tom
	fmt.Printf("p_person: %v\n", p_person)
	fmt.Printf("*p_person: %v\n", *p_person)

	// 使用new 关键字创建结构体指针
	var p_person2 = new(Person)
	p_person2 = &tom
	fmt.Printf("p_person2: %v\n", p_person2)
	fmt.Println((*p_person).age)
	fmt.Println(p_person.age)
}
