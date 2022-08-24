package main

import (
	"fmt"
	"time"
)

var c1 = make(chan int)
var c2 = make(chan string)

func main() {
	go func() {
		c1 <- 100
		c2 <- "hello"
		close(c1)
		close(c2)
	}()

	for {
		select {
		case r := <-c1:
			fmt.Print(r)
		case r := <-c2:
			fmt.Print(r)
		default:
			fmt.Print("default")
		}
		time.Sleep(time.Second)
	}
}
