package main

import "fmt"

func main() {
	// 类型定义
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i: %v\n", i)

	// 类型别名
	type MyInt2 = int
	var i1 MyInt2
	i1 = 1000
	fmt.Printf("i1: %v\n", i1)
}
