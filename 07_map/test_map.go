package main

import "fmt"

func main() {
	var m1 = make(map[string]string)
	m1["name"] = "tom"
	fmt.Printf("m1[\"name\"]: %v\n", m1["name"])

	var m2 = map[string]string{"name": "tom"}
	fmt.Printf("m2[\"name\"]: %v\n", m2["name"])

	// 判断key是否存在
	v, ok := m1["age"]
	v1, ok1 := m1["name"]
	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("ok1: %v\n", ok1)
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
}
