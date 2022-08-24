package main

import (
	"fmt"
	"time"
)

func show_msg(msg string) {
	for i := 0; i < 4; i++ {
		fmt.Printf("%v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go show_msg("java") // 开启一个协程
	show_msg("go")
}
