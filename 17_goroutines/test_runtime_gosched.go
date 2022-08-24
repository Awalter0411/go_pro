package main

import (
	"fmt"
	"runtime"
)

func showMsg(s string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("s: %v\n", s)
	}
}

func main() {
	go showMsg("java")
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 让其他子协程来执行
		fmt.Printf("%v\n", i)
	}
	fmt.Print("end...")
}
