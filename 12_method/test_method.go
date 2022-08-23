package main

import "fmt"

type Person struct {
	name string
}

type User struct {
	username string
	password string
}

// recevier
func (per Person) eat() {
	fmt.Printf("%v eat\n", per.name)
}

func (per Person) sleep() {
	fmt.Printf("%v sleep\n", per.name)
}

func (user User) isLogin(username string, age string) bool {
	if user.username == "tom" && user.password == "123" {
		return true
	} else {
		return false
	}
}

func main() {
	var p = Person{name: "curry"}
	p.eat()
	p.sleep()
	var user = User{
		username: "tom",
		password: "123",
	}
	user.isLogin("tom", "123")
}
