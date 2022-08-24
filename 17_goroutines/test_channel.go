package main

import (
	"fmt"
	"time"
)

// 无缓冲通道
var values = make(chan int)

func send() {
	value := 10
	fmt.Printf("send value: %v\n", value)
	time.Sleep(time.Second * 3)
	values <- value
}

func main() {
	go send()
	value := <-values
	fmt.Printf("get value: %v\n", value)
}
