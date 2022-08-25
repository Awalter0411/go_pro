package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func main() {
	for j := 0; j < 100; j++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("i: %v\n", i)
}
