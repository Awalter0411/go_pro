package main

import "fmt"

var c = make(chan int)

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		// 需要关闭，否则可能会出现死锁
		close(c)
	}()

	// 第一种遍历方式
	// for i := 0; i < 3; i++ {
	// 	r := <-c
	// 	fmt.Printf("r: %v\n", r)
	// }

	// 第二种遍历方式
	// for v := range c {
	// 	fmt.Printf("v: %v\n", v)
	// }

	// 第三种
	for {
		v, ok := <-c
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			break
		}
	}
}
