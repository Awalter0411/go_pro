package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	name := "tom"
	age := "20"
	s := strings.Join([]string{name, age}, "-")
	fmt.Printf("s: %v\n", s)

	// 使用buffer
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer: %v\n", buffer.String())

	// 字符串切片
	s1 := "hello world"
	a := 2
	b := 5
	fmt.Printf("s1[a:b]: %v\n", s1[a:b])
	fmt.Printf("s1[a:]: %v\n", s1[a:])
	fmt.Printf("s1[:b]: %v\n", s1[:b])

	// 字符串函数
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))
}
