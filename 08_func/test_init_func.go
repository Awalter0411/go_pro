package main

import "fmt"

var s = demo()

func init() {
	fmt.Println("init")
}

func demo() string {
	fmt.Println("init var")
	return "hello"
}

func main() {
	// 初始化顺序：变量初始化->init()->main()
	fmt.Println("main")
}
