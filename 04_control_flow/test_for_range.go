package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("index is %v ", i)
		fmt.Printf("value is %v\n", v)
	}
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	for k, v := range m {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}
