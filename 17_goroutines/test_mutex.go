package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup

// 互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	i++
	fmt.Printf("i: %v\n", i)
	lock.Unlock()
}

func sub() {
	defer wg.Done()
	lock.Lock()
	i--
	fmt.Printf("i: %v\n", i)
	lock.Unlock()
}

func main() {
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Printf("end i: %v\n", i)
}
