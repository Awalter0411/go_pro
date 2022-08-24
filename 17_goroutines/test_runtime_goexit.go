package main

import (
	"fmt"
	"runtime"
	"time"
)

func showMsg() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			runtime.Goexit()
		}
		fmt.Printf("i:%v\n", i)
	}
}

func main() {
	go showMsg()
	time.Sleep(time.Second)
}
