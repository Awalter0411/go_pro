package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a i: %v\n", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b i: %v\n", i)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	// 设置cpu核心数
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}
