package main

import "fmt"

type Animal struct {
	name string
	age  int
}

type Dog struct {
	a Animal
}

func (d Dog) wang() {
	fmt.Println("wang")
}

type Cat struct {
	Animal
}

func (c Cat) miao() {
	fmt.Println("miao")
}

func main() {
	var d = Dog{
		a: Animal{
			name: "jame",
			age:  30,
		},
	}
	d.wang()
	fmt.Printf("d: %v\n", d)

	var c = Cat{
		Animal{
			name: "jame",
			age:  30,
		},
	}
	fmt.Printf("c.name: %v\n", c.name)
}
