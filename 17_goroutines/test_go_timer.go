package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1: %v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2: %v\n", t2)

	// 阻塞
	<-time.After(time.Second * 2)
	fmt.Print("2s 后")

	// 停止
	stop := timer1.Stop()
	if stop {
		fmt.Print("123")
	}

	// 重置
	timer2 := time.NewTimer(time.Second)
	timer2.Reset(time.Second * 2)
}
