package main

import "fmt"

type Person struct {
	name string
	age  int
}

func new_person(name string, age int) (*Person, error) {
	if name == "" || age < 0 {
		return nil, fmt.Errorf("name 不能为空, age不能小于0")
	}
	return &Person{name: name, age: age}, nil
}

func main() {
	p, err := new_person("tom", 30)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("err: %v\n", err)
}
